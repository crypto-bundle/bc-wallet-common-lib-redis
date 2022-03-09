package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type MockSrv struct {
	t *testing.T
}

func (m *MockSrv) NewLoggerEntry(_ string) (*zap.Logger, error) {
	return zaptest.NewLogger(m.t), nil
}

func (m *MockSrv) NewLoggerEntryWithFields(named string, fields ...zap.Field) (*zap.Logger, error) {
	return zaptest.NewLogger(m.t), nil
}

func NewLoggerServiceMock(t *testing.T) *MockSrv {
	return &MockSrv{t: t}
}
