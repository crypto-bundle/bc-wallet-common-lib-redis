package metric

import "errors"

var (
	ErrPanicRecovered     = errors.New("panic recovered")
	ErrMetricNotFound     = errors.New("metric not found")
	ErrMismatchMetricType = errors.New("mismatch metric type")
	ErrMetricCmdNotFound  = errors.New("metric command not found")
)
