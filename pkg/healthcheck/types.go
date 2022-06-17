package healthcheck

type config interface {
	IsDebug() bool
	GetAddress() string
	GetLivenessPath() string
	GetReadinessPath() string
	GetStartupPath() string
}
