package healthcheck

import (
	"context"
	"net/http"
	"time"
)

type unitParamsService interface {
	IsDebug() bool
	GetProbeName() string

	GetHTTPListenAddress() string
	GetHTTPHandlerPath() string
	GetHTTPReadTimeout() time.Duration
	GetHTTPWriteTimeout() time.Duration
}

type configService interface {
	GetLivenessParams() unitParamsService
	GetReadinessParams() unitParamsService
	GetStartupParams() unitParamsService
}

type probeService interface {
	Do(ctx context.Context) Status
	Init(ctx context.Context) error
	Shutdown(ctx context.Context) error
	ListenAndServe(ctx context.Context) error
}

type handlerService interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type healthCheckerService interface {
	Init() error
}

type appDirectiveExecutionService interface {
	Do(status Status) (httpStatus int, httpBody string)
}
