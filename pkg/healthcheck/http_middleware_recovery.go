package healthcheck

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var (
	ErrHealthCheckRecovery = errors.New("healthcheck recovery from panic")
)

type middleware struct {
	logger      *zap.Logger
	httpHandler http.Handler
}

type middlewareRecovery struct {
	logger *zap.Logger
}

func newRecoveryMiddleware(l *zap.Logger) *middlewareRecovery {
	return &middlewareRecovery{
		logger: l,
	}
}

func (m *middlewareRecovery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Add("Content-Type", "text/plain")
			respText := fmt.Sprintf("%e\n%+v\n", ErrHealthCheckRecovery, recoverErr)
			_, writeErr := w.Write([]byte(respText))
			if writeErr != nil {
				m.logger.Error(
					"unable to write response",
					zap.Error(writeErr),
					zap.Time(RecoveryTimeTag, time.Now()),
				)
			}

			m.logger.Error(
				ErrHealthCheckRecovery.Error(),
				zap.Any(RecoveryErrTag, recoverErr),
				zap.Time(RecoveryTimeTag, time.Now()),
				zap.Stack(RecoveryStackTag),
			)
		}
	}()

	return
}

func (m *middleware) With(next http.Handler) *middleware {
	m.httpHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})

	return m
}

func (m *middleware) GetHTTPHandler() http.Handler {
	return m.httpHandler
}

func newMiddleware(l *zap.Logger) *middleware {
	return &middleware{logger: l}
}
