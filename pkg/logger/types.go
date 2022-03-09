package logger

type configManager interface {
	IsDebug() bool
	GetMinimalLogLevel() string
}
