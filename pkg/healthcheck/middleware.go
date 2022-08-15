package healthcheck

import (
	"net/http"
	"time"
	"go.uber.org/zap"
	"runtime/debug"
)

type middleware struct {
	logger *zap.Logger
}

func NewMiddleware(logger *zap.Logger) *middleware {
	return &middleware{logger: logger}
}

func (m *middleware) Wrap(next http.Handler, debug bool) http.Handler {
	if debug {
		next = m.withDebugInfo(next)
	}

	return m.withRecovery(next)
}

func (m *middleware) withRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				m.logger.Error(
					"[Recovery from panic]",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (m *middleware) withDebugInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			m.logger.Info(
				r.URL.Path,
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("query", r.URL.RawQuery),
				zap.String("ip", r.RemoteAddr),
				zap.String("user-agent", r.UserAgent()),
				zap.Duration("latency", time.Now().Sub(start)),
			)
		}()

		next.ServeHTTP(w, r)
	})
}
