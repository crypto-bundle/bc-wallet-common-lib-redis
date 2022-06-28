package healthcheck

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServiceCreator struct {
	cfg    config
	logger *zap.Logger
}

func NewServiceCreator(cfg config, logger *zap.Logger) *ServiceCreator {
	return &ServiceCreator{cfg: cfg, logger: logger}
}

func (c *ServiceCreator) Create(livenessProbe, readinessProbe, startupProbe Probe) (HealthcheckService, error) {
	healthcheckSrv := New(c.cfg, livenessProbe, readinessProbe, startupProbe, c.logger)
	err := healthcheckSrv.Init()
	if err != nil {
		return nil, err
	}

	return healthcheckSrv, nil
}

type Service struct {
	cfg    config
	logger *zap.Logger

	handler *httpHandler
}

func New(cfg config, livenessProbe, readinessProbe, startupProbe Probe, logger *zap.Logger) *Service {
	l := logger.Named("healthcheck")

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
			s.logger.Error("unable to listen and serve http server", zap.Error(err))
		}
	}()

	s.logger.Info("initiated successfully", zap.String("address", s.cfg.GetAddress()))

	return nil
}
