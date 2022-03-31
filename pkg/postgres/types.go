package postgres

type DbConfig interface {
	GetDbHost() string
	GetDbPort() uint16
	GetDbName() string
	GetDbUser() string
	GetDbPassword() string
	GetDbTLSMode() string
	GetDbRetryCount() uint8
	GetDbConnectTimeOut() uint16

	GetDbMaxOpenConns() uint8
	GetDbMaxIdleConns() uint8

	IsDebug() bool
}
