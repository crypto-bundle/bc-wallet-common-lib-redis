package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

type ConnectionParams struct {
	redisConfigService

	sslMode string

	debug bool
}

type Connection struct {
	params *ConnectionParams
	logger *zap.Logger

	client *redis.Client
}

func (c *Connection) IsHealed(ctx context.Context) bool {
	statusCmd := c.client.Ping(ctx)

	res, err := statusCmd.Result()
	if err != nil {
		return false
	}

	return res == "PONG"
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
	retryCount := c.params.GetRetryConnCount()
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
			time.Sleep(c.params.GetRetryConnTimeOut())

			continue
		}

		_, err = client.Ping(ctx).Result()
		if err != nil {
			c.logger.Error("unable ping redis. reconnecting...",
				zap.Error(err), zap.Int("iteration", try),
				zap.Any("params", c.params))
			try++
			time.Sleep(c.params.GetRetryConnTimeOut())

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
		MaxRetries:         int(params.GetMaxRetryCount()),
		MinRetryBackoff:    -1,
		MaxRetryBackoff:    -1,
		DialTimeout:        params.GetDialTimeout(),
		ReadTimeout:        params.GetReadTimeOut(),
		WriteTimeout:       params.GetWriteTimeOut(),
		PoolFIFO:           true,
		PoolSize:           int(params.GetPoolSize()),
		MinIdleConns:       int(params.GetMinIdleConn()),
		MaxConnAge:         params.GetMaxConnectionAge(),
		PoolTimeout:        params.GetPoolTimeout(),
		IdleTimeout:        params.GetIdleTimeout(),
		IdleCheckFrequency: time.Second * 60,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	return client, nil
}

// NewConnection to redis server
func NewConnection(_ context.Context, cfg redisConfigService, logger *zap.Logger) *Connection {
	conn := &Connection{
		params: &ConnectionParams{
			redisConfigService: cfg,
			sslMode:            "",
			debug:              false,
		},
		logger: logger,
		client: nil,
	}

	return conn
}
