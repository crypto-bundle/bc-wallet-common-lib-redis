package healthcheck

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Probe func(ctx context.Context) error

type Service struct {
	cfg    config
	logger *zap.Logger

	handler *httpHandler
}

func New(cfg config, livenessProbe, readinessProbe, startupProbe Probe, loggerSrv *zap.Logger) *Service {
	l := loggerSrv.Named("healthcheck")

	h := NewHttpHandler(livenessProbe, readinessProbe, startupProbe)

	return &Service{
		cfg:     cfg,
		logger:  l,
		handler: h,
	}
}

func (s *Service) Init() error {
	if !s.cfg.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET(s.cfg.GetLivenessPath(), s.handler.Liveness)
	router.GET(s.cfg.GetReadinessPath(), s.handler.Readiness)
	router.GET(s.cfg.GetStartupPath(), s.handler.Startup)

	server := &http.Server{
		Addr:         s.cfg.GetAddress(),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			s.logger.Error("unable to listen and serve healthcheck http server", zap.Error(err))
		}
	}()

	return nil
}
