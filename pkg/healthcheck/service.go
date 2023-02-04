package healthcheck

import (
	"net/http"
	"time"

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
	mux := http.NewServeMux()
	mux.HandleFunc(s.cfg.GetLivenessPath(), s.handler.livenessHandler)
	mux.HandleFunc(s.cfg.GetReadinessPath(), s.handler.getReadinessHandler)
	mux.HandleFunc(s.cfg.GetStartupPath(), s.handler.getStartupHandler)

	middleware := NewMiddleware(s.logger)

	server := &http.Server{
		Addr:         s.cfg.GetAddress(),
		Handler:      middleware.Wrap(mux, s.cfg.IsDebug()),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     zap.NewStdLog(s.logger),
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
