package router

import "go.uber.org/zap"

type Service struct {
	l *zap.Logger

	routes map[int]map[int]struct{}
}
