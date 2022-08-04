package metric

import (
	"time"
)

const (
	DefaultHttpReadTimeout  = 5 * time.Second
	DefaultHttpWriteTimeout = 10 * time.Second
)

const (
	LoggerTagRecover = "recover"
	LoggerTagMetric  = "metric"
	LoggerTagCmd     = "command"
)
