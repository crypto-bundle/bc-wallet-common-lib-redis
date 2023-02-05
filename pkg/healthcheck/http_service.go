package healthcheck

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

var (
	ErrProbeAlreadyExist = errors.New("passed probe already exist")
)

type httpHealthChecker struct {
	logger *zap.Logger

	probes [3]probeService // liveness, rediness, startup
}

func (s *httpHealthChecker) Init(ctx context.Context) error {
	for _, probe := range s.probes {
		if probe == nil {
			continue
		}

		err := probe.Init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *httpHealthChecker) ListenAndServe(ctx context.Context) error {
	for _, probe := range s.probes {
		if probe == nil {
			continue
		}

		go func(probeSrv probeService) {
			err := probeSrv.ListenAndServe(ctx)
			if err != nil {
				s.logger.Info("unable to start listen and server process for probe")
			}
		}(probe)
	}

	s.logger.Info("all probes successfully listen up")

	return nil
}

func (s *httpHealthChecker) AddLivenessProbe(probe probeService) error {
	if s.probes[LivenessProbeIndex] != nil {
		return ErrProbeAlreadyExist
	}

	s.probes[LivenessProbeIndex] = probe

	return nil
}

func (s *httpHealthChecker) AddRedinessProbe(probe probeService) error {
	if s.probes[RedinessProbeIndex] != nil {
		return ErrProbeAlreadyExist
	}

	s.probes[RedinessProbeIndex] = probe

	return nil
}

func (s *httpHealthChecker) AddStartupProbe(probe probeService) error {
	if s.probes[StartupProbeIndex] != nil {
		return ErrProbeAlreadyExist
	}

	s.probes[StartupProbeIndex] = probe

	return nil
}

func (s *httpHealthChecker) Shutdown(ctx context.Context) error {
	for _, probe := range s.probes {
		if probe == nil {
			continue
		}

		go func(probeSrv probeService) {
			err := probeSrv.Shutdown(ctx)
			if err != nil {
				s.logger.Info("unable to stop probe")
			}
		}(probe)
	}

	s.logger.Info("all probes successfully stopped")

	return nil
}

func NewHTTPHealthChecker(l *zap.Logger) *httpHealthChecker {
	healthChecker := &httpHealthChecker{
		logger: l,
		probes: [3]probeService{},
	}

	return healthChecker
}
