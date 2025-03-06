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
	"fmt"
	"time"
)

type RedisConfig struct {
	Host     string `envconfig:"REDIS_SERVICE_HOST" default:"redis.local"`
	Port     uint16 `envconfig:"REDIS_SERVICE_PORT" default:"6380"`
	User     string `envconfig:"REDIS_USER" secret:"true" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" secret:"true" required:"true"`
	Database int    `envconfig:"REDIS_DB" default:"1"`
	// RetryConnTimeOut is the maximum number of reconnection tries. If 0 - infinite loop
	RetryConnTimeOut time.Duration `envconfig:"REDIS_CONNECTION_RETRY_TIMEOUT" default:"1s"`
	// RetryConnCount is the timeout in millisecond to connect between connection tries
	RetryConnCount uint8 `envconfig:"REDIS_CONNECTION_RETRY_COUNT" default:"0"`
	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetryCount uint8 `envconfig:"REDIS_MAX_RETRY_COUNT" default:"3"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeOut time.Duration `envconfig:"REDIS_READ_TIMEOUT" default:"3s"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeOut time.Duration `envconfig:"REDIS_WRITE_TIMEOUT" default:"3s"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConn uint8 `envconfig:"REDIS_MIN_IDLE_CONNECTIONS" default:"0"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration `envconfig:"REDIS_IDLE_TIMEOUT" default:"5m"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnectionAge time.Duration `envconfig:"REDIS_MAX_CONNECTION_AGE" default:"0"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `envconfig:"REDIS_POOL_TIMEOUT" default:"4s"`
	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize uint8 `envconfig:"REDIS_POOL_SIZE" default:"10"`
	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `envconfig:"REDIS_DIAL_TIMEOUT" default:"5s"`
}

// Prepare variables to static configuration
func (c *RedisConfig) Prepare() error {
	return nil
}

func (c *RedisConfig) PrepareWith(cfgSrvList ...interface{}) error {
	return nil
}

func (c *RedisConfig) GetRedisHost() string {
	return c.Host
}

func (c *RedisConfig) GetRedisPort() uint16 {
	return c.Port
}

func (c *RedisConfig) GetRedisAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *RedisConfig) GetRedisDbName() int {
	return c.Database
}

func (c *RedisConfig) GetRedisUser() string {
	return c.User
}

func (c *RedisConfig) GetRedisPassword() string {
	return c.Password
}

func (c *RedisConfig) GetRetryConnCount() uint8 {
	return c.RetryConnCount
}

func (c *RedisConfig) GetRetryConnTimeOut() time.Duration {
	return c.RetryConnTimeOut
}

func (c *RedisConfig) GetMaxRetryCount() uint8 {
	return c.MaxRetryCount
}

func (c *RedisConfig) GetDialTimeout() time.Duration {
	return c.DialTimeout
}

func (c *RedisConfig) GetReadTimeOut() time.Duration {
	return c.ReadTimeOut
}

func (c *RedisConfig) GetPoolTimeout() time.Duration {
	return c.PoolTimeout
}

func (c *RedisConfig) GetIdleTimeout() time.Duration {
	return c.IdleTimeout
}

func (c *RedisConfig) GetWriteTimeOut() time.Duration {
	return c.WriteTimeOut
}

func (c *RedisConfig) GetPoolSize() uint8 {
	return c.PoolSize
}

func (c *RedisConfig) GetMinIdleConn() uint8 {
	return c.MinIdleConn
}

func (c *RedisConfig) GetMaxConnectionAge() time.Duration {
	return c.MaxConnectionAge
}
