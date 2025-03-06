/*
 * MIT NON-AI License
 *
 * Copyright (c) 2022-2025 Aleksei Kotelnikov(gudron2s@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of the software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions.
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * In addition, the following restrictions apply:
 *
 * 1. The Software and any modifications made to it may not be used for the purpose of training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining. This condition applies to any derivatives,
 * modifications, or updates based on the Software code. Any usage of the Software in an AI-training dataset is considered a breach of this License.
 *
 * 2. The Software may not be included in any dataset used for training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining.
 *
 * 3. Any person or organization found to be in violation of these restrictions will be subject to legal action and may be held liable
 * for any damages resulting from such use.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-redis/redis/v8"
)

type ConnectionParams struct {
	redisConfigService

	sslMode string

	debug bool
}

type Connection struct {
	l *slog.Logger
	e errorFormatterService

	params *ConnectionParams

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
		return c.e.ErrorOnly(err)
	}

	return nil
}

// Connect to redis cache storage...
func (c *Connection) Connect(ctx context.Context) (*Connection, error) {
	retryDecValue := uint8(1)
	retryCount := c.params.GetRetryConnCount()

	if retryCount == 0 {
		retryDecValue = 0
		retryCount = 1
	}

	try := 0

	for i := retryCount; i != 0; i -= retryDecValue {
		client := prepareClient(c.params)

		_, err := client.Ping(ctx).Result()
		if err != nil {
			c.l.Error("unable ping redis. reconnecting...", slog.Any("error", err),
				slog.Int("iteration", try),
				slog.Any("params", c.params))

			try++

			time.Sleep(c.params.GetRetryConnTimeOut())

			continue
		}

		c.client = client

		return c, nil
	}

	return c, nil
}

func prepareClient(params *ConnectionParams) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:               params.GetRedisAddress(),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           params.GetRedisUser(),
		Password:           params.GetRedisPassword(),
		DB:                 params.GetRedisDBName(),
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
		IdleCheckFrequency: time.Minute,
		TLSConfig:          nil,
		Limiter:            nil,
		Network:            "",
	})

	return client
}

// NewConnection to redis server...
func NewConnection(logBuilder loggerFabricService,
	errFmtSvc errorFormatterService,
	cfg redisConfigService,
) *Connection {
	conn := &Connection{
		l: logBuilder.NewSlogNamedLoggerEntry(redisNameSpace),
		e: errFmtSvc,
		params: &ConnectionParams{
			redisConfigService: cfg,
			sslMode:            "",
			debug:              false,
		},
		client: nil,
	}

	return conn
}
