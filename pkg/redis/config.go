package redis

import (
	"fmt"
	"time"
)

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" json:"-"`
	Port     uint16 `envconfig:"REDIS_PORT" json:"-"`
	User     string `envconfig:"REDIS_USER" json:"REDIS_USER"`
	Password string `envconfig:"REDIS_PASSWORD" json:"REDIS_PASSWORD"`
	Database int    `envconfig:"REDIS_DB" json:"REDIS_DB"`
	// RetryConnTimeOut is the maximum number of reconnection tries. If 0 - infinite loop
	RetryConnTimeOut time.Duration `envconfig:"REDIS_CONNECTION_RETRY_TIMEOUT" default:"1s" json:"-"`
	// RetryConnCount is the timeout in millisecond to connect between connection tries
	RetryConnCount uint8 `envconfig:"REDIS_CONNECTION_RETRY_COUNT" default:"0" json:"-"`
	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetryCount uint8 `envconfig:"REDIS_MAX_RETRY_COUNT" default:"3" json:"-"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeOut time.Duration `envconfig:"REDIS_READ_TIMEOUT" default:"3s" json:"-"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeOut time.Duration `envconfig:"REDIS_WRITE_TIMEOUT" default:"3s" json:"-"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConn uint8 `envconfig:"REDIS_MIN_IDLE_CONNECTIONS" json:"-"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration `envconfig:"REDIS_IDLE_TIMEOUT" default:"5m" json:"-"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnectionAge time.Duration `envconfig:"REDIS_MAX_CONNECTION_AGE" json:"-"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `envconfig:"REDIS_POOL_TIMEOUT" default:"4s" json:"-"`
	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize uint8 `envconfig:"REDIS_POOL_SIZE" default:"10" json:"-"`
	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `envconfig:"REDIS_DIAL_TIMEOUT" default:"5s" json:"-"`
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
