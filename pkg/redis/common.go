package redis

import "time"

type redisConfigService interface {
	GetRedisHost() string
	GetRedisPort() uint16
	GetRedisAddress() string
	GetRedisDbName() int
	GetRedisUser() string
	GetRedisPassword() string

	GetRetryConnCount() uint8
	GetRetryConnTimeOut() time.Duration
	GetMaxRetryCount() uint8

	GetDialTimeout() time.Duration
	GetReadTimeOut() time.Duration
	GetWriteTimeOut() time.Duration
	GetPoolTimeout() time.Duration
	GetIdleTimeout() time.Duration

	GetPoolSize() uint8

	GetMinIdleConn() uint8
	GetMaxConnectionAge() time.Duration
}
