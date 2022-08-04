package metric

import "errors"

var (
	ErrMetricNotFound     = errors.New("metric not found")
	ErrMismatchMetricType = errors.New("mismatch metric type")
	ErrMetricCmdNotFound  = errors.New("metric command not found")
)
