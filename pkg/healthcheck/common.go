package healthcheck

import (
	"context"
	"net/http"
	"time"
)

type unitParamsService interface {
	IsDebug() bool
	GetUnitName() string

	GetHTTPListenAddress() string
	GetHTTPHandlerPath() string
	GetHTTPReadTimeout() time.Duration
	GetHTTPWriteTimeout() time.Duration
}

type probeService interface {
	Do(ctx context.Context) Status
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
