package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

type ConnectionParams struct {
	*RedisConfig

	sslMode string

	debug bool
}

type Connection struct {
	params *ConnectionParams
	logger *zap.Logger

	client *redis.Client
}

func (c *Connection) GetClient() *redis.Client {
	return c.client
}

func (c *Connection) Close() error {
	err := c.client.Close()
	if err != nil {
		return err
	}

	return nil
}

// Connect to postgres database
func (c *Connection) Connect(ctx context.Context) (*Connection, error) {
	retryDecValue := uint8(1)
	retryCount := c.params.RetryConnCount
	if retryCount == 0 {
		retryDecValue = 0
		retryCount = 1
	}
	try := 0

	for i := retryCount; i != 0; i -= retryDecValue {
		client, err := prepareClient(c.params)
		if err != nil {
			c.logger.Error("unable to prepare redis client. reconnecting...",
				zap.Error(err), zap.Int("iteration", try))
			try++
			time.Sleep(c.params.RetryConnTimeOut)

			continue
		}

		_, err = client.Ping(ctx).Result()
		if err != nil {
			c.logger.Error("unable ping redis. reconnecting...",
				zap.Error(err), zap.Int("iteration", try),
				zap.Any("params", c.params))
			try++
			time.Sleep(c.params.RetryConnTimeOut)

			continue
		}

		c.client = client
		return c, nil
	}

	return c, nil
}

func prepareClient(params *ConnectionParams) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:               params.GetRedisAddress(),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           params.GetRedisUser(),
		Password:           params.GetRedisPassword(),
		DB:                 params.GetRedisDbName(),
		MaxRetries:         int(params.MaxRetryCount),
		MinRetryBackoff:    -1,
		MaxRetryBackoff:    -1,
		DialTimeout:        params.DialTimeout,
		ReadTimeout:        params.ReadTimeOut,
		WriteTimeout:       params.WriteTimeOut,
		PoolFIFO:           true,
		PoolSize:           int(params.PoolSize),
		MinIdleConns:       int(params.MinIdleConn),
		MaxConnAge:         params.MaxConnectionAge,
		PoolTimeout:        params.PoolTimeout,
		IdleTimeout:        params.IdleTimeout,
		IdleCheckFrequency: time.Second * 60,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	return client, nil
}

// NewConnection to redis server
func NewConnection(ctx context.Context, cfg *RedisConfig, logger *zap.Logger) *Connection {
	conn := &Connection{
		params: &ConnectionParams{
			RedisConfig: cfg,
			sslMode:     "",
			debug:       false,
		},
		logger: logger,
		client: nil,
	}

	return conn
}
