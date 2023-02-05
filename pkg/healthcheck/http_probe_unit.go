package healthcheck

import (
	"context"
	"net/http"
	"syscall"

	"go.uber.org/zap"
)

type probeUnit struct {
	logger *zap.Logger
	cfg    unitParamsService

	requestHandler http.Handler

	httpSrv *http.Server

	applicationPID int
}

func (s *probeUnit) Init() error {
	s.applicationPID = syscall.Getpid()

	mux := http.NewServeMux()

	httpMiddleware := newMiddleware(s.logger)
	handlerWithMiddleware := httpMiddleware.With(s.requestHandler).
		With(newRecoveryMiddleware(s.logger))

	mux.Handle(s.cfg.GetHTTPHandlerPath(), handlerWithMiddleware.GetHTTPHandler())

	server := &http.Server{
		Addr:         s.cfg.GetHTTPListenAddress(),
		Handler:      mux,
		ReadTimeout:  s.cfg.GetHTTPReadTimeout(),
		WriteTimeout: s.cfg.GetHTTPWriteTimeout(),
		ErrorLog:     zap.NewStdLog(s.logger),
	}

	s.httpSrv = server

	s.logger.Info("initiated successfully")

	return nil
}

func (s *probeUnit) ListenAndServe() error {
	err := s.httpSrv.ListenAndServe()
	if err != nil {
		s.logger.Error("unable to listen and serve http server", zap.Error(err))
		return err
	}

	s.logger.Info("run successfully")

	return nil
}

func (s *probeUnit) Shutdown(ctx context.Context) error {
	err := s.httpSrv.Shutdown(ctx)
	if err != nil {
		s.logger.Error("unable to stop http server", zap.Error(err))
		return err
	}

	return nil
}

func NewHTPPHealthCheckerUnit(paramsSrv unitParamsService,
	logger *zap.Logger,
	probeSrv probeService,
) *probeUnit {
	l := logger.Named("healthcheck_unit").
		With(zap.String(ListenAddressTag, paramsSrv.GetHTTPListenAddress())).
		With(zap.String(UnitNameTag, paramsSrv.GetProbeName()))

	return &probeUnit{
		cfg:            paramsSrv,
		logger:         l,
		applicationPID: -1,

		requestHandler: newHttpHandler(probeSrv),
	}
}
