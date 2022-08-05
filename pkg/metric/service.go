package metric

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	ginzap "github.com/gin-contrib/zap"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	mu sync.Mutex

	name   string
	cfg    config
	logger *zap.Logger

	metrics map[string]prometheus.Collector

	cmdToCountersCh     chan cmdForCounter
	cmdToCounterVecsCh  chan cmdForCounterVec
	cmdToGaugesCh       chan cmdForGauge
	cmdToGaugeVecsCh    chan cmdForGaugeVec
	cmdToObserversCh    chan cmdForObserver
	cmdToObserverVecsCh chan cmdForObserverVec
}

func New(name string, cfg config, logger *zap.Logger) *Service {
	return &Service{
		name:                name,
		cfg:                 cfg,
		logger:              logger.Named("metric"),
		cmdToCountersCh:     make(chan cmdForCounter, 1),
		cmdToCounterVecsCh:  make(chan cmdForCounterVec, 1),
		cmdToGaugesCh:       make(chan cmdForGauge, 1),
		cmdToGaugeVecsCh:    make(chan cmdForGaugeVec, 1),
		cmdToObserversCh:    make(chan cmdForObserver, 1),
		cmdToObserverVecsCh: make(chan cmdForObserverVec, 1),
	}
}

func (s *Service) Init() error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		ginzap.Ginzap(s.logger, time.RFC3339, false),
		ginzap.RecoveryWithZap(s.logger, true),
	)
	router.GET(s.cfg.GetPath(), gin.WrapH(promhttp.Handler()))

	server := &http.Server{
		Addr:         s.cfg.GetAddress(),
		Handler:      router,
		ReadTimeout:  DefaultHttpReadTimeout,
		WriteTimeout: DefaultHttpWriteTimeout,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			s.logger.Error("unable to listen and serve http server", zap.Error(err))
		}
	}()

	go s.cmdLoop()

	s.logger.Info("initiated successfully", zap.String("address", s.cfg.GetAddress()))

	return nil
}

func (s *Service) AddCounter(name, help string) {
	c := promauto.NewCounter(prometheus.CounterOpts{
		Name: s.formatMetricName(name),
		Help: help,
	})

	s.addMetric(name, c)
}

func (s *Service) AddCounterVec(name, help string, labelNames []string) {
	cv := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: s.formatMetricName(name),
		Help: help,
	}, labelNames)

	s.addMetric(name, cv)
}

func (s *Service) AddCounterFunc(name, help string, callback func() float64) {
	cf := promauto.NewCounterFunc(prometheus.CounterOpts{
		Name: s.formatMetricName(name),
		Help: help,
	}, callback)

	s.addMetric(name, cf)
}

func (s *Service) AddGauge(name, help string) {
	g := promauto.NewGauge(prometheus.GaugeOpts{
		Name: s.formatMetricName(name),
		Help: help,
	})

	s.addMetric(name, g)
}

func (s *Service) AddGaugeVec(name, help string, labelNames []string) {
	gv := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: s.formatMetricName(name),
		Help: help,
	}, labelNames)

	s.addMetric(name, gv)
}

func (s *Service) AddGaugeFunc(name, help string, callback func() float64) {
	gf := promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: s.formatMetricName(name),
		Help: help,
	}, callback)

	s.addMetric(name, gf)
}

func (s *Service) AddHistogram(name, help string, buckets []float64) {
	h := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    s.formatMetricName(name),
		Help:    help,
		Buckets: buckets,
	})

	s.addMetric(name, h)
}

func (s *Service) AddHistogramVec(name, help string, labelNames []string, buckets []float64) {
	hv := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    s.formatMetricName(name),
		Help:    help,
		Buckets: buckets,
	}, labelNames)

	s.addMetric(name, hv)
}

func (s *Service) AddSummary(name, help string) {
	summary := promauto.NewSummary(prometheus.SummaryOpts{
		Name: s.formatMetricName(name),
		Help: help,
	})

	s.addMetric(name, summary)
}

func (s *Service) AddSummaryVec(name, help string, labelNames []string) {
	sv := promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: s.formatMetricName(name),
		Help: help,
	}, labelNames)

	s.addMetric(name, sv)
}

func (s *Service) GetCounter(name string) (prometheus.Counter, error) {
	metric, ok := s.metrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	counter, ok := metric.(prometheus.Counter)
	if !ok {
		return nil, ErrMismatchMetricType
	}

	return counter, nil
}

func (s *Service) GetCounterVec(name string) (*prometheus.CounterVec, error) {
	metric, ok := s.metrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	vec, ok := metric.(*prometheus.CounterVec)
	if !ok {
		return nil, ErrMismatchMetricType
	}

	return vec, nil
}

func (s *Service) GetGauge(name string) (prometheus.Gauge, error) {
	metric, ok := s.metrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	gauge, ok := metric.(prometheus.Gauge)
	if !ok {
		return nil, ErrMismatchMetricType
	}

	return gauge, nil
}

func (s *Service) GetGaugeVec(name string) (*prometheus.GaugeVec, error) {
	metric, ok := s.metrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	vec, ok := metric.(*prometheus.GaugeVec)
	if !ok {
		return nil, ErrMismatchMetricType
	}

	return vec, nil
}

func (s *Service) GetObserver(name string) (prometheus.Observer, error) {
	metric, ok := s.metrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	observer, ok := metric.(prometheus.Observer)
	if !ok {
		return nil, ErrMismatchMetricType
	}

	return observer, nil
}

func (s *Service) GetObserverVec(name string) (prometheus.ObserverVec, error) {
	metric, ok := s.metrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	vec, ok := metric.(prometheus.ObserverVec)
	if !ok {
		return nil, ErrMismatchMetricType
	}

	return vec, nil
}

func (s *Service) formatMetricName(name string) string {
	return s.name + "_" + strings.ReplaceAll(name, "-", "_")
}

func (s *Service) addMetric(name string, metric prometheus.Collector) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.metrics == nil {
		s.metrics = map[string]prometheus.Collector{}
	}

	s.metrics[name] = metric
}

func (s *Service) cmdLoop() {
	defer func() {
		s.logger.Error("panic recovered", zap.Any(LoggerTagRecover, recover()))

		s.cmdLoop()
	}()
	for {
		select {
		case e := <-s.cmdToCountersCh:
			s.handleCounterCmd(e)
		case e := <-s.cmdToCounterVecsCh:
			s.handleCounterVecCmd(e)
		case e := <-s.cmdToGaugesCh:
			s.handleGaugeCmd(e)
		case e := <-s.cmdToGaugeVecsCh:
			s.handleGaugeVecCmd(e)
		case e := <-s.cmdToObserversCh:
			s.handleObserverCmd(e)
		case e := <-s.cmdToObserverVecsCh:
			s.handleObserverVecCmd(e)
		}
	}
}

func (s *Service) handleCounterCmd(cmd cmdForCounter) {
	m, err := s.GetCounter(cmd.GetMetricName())
	if err != nil {
		s.logger.Error(
			"unable to get the counter metric",
			zap.Error(err),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
		)
		return
	}

	switch cmd.GetType() {
	case CmdCounterInc:
		m.Inc()
	case CmdCounterAdd:
		m.Add(cmd.GetValue())
	default:
		s.logger.Error(
			"unable to find requested counter metric command",
			zap.Error(ErrMetricCmdNotFound),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
			zap.String(LoggerTagCmd, cmd.GetType().String()),
		)
	}
}

func (s *Service) handleCounterVecCmd(cmd cmdForCounterVec) {
	m, err := s.GetCounterVec(cmd.GetMetricName())
	if err != nil {
		s.logger.Error(
			"unable to get the counter vec metric",
			zap.Error(err),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
		)
		return
	}

	switch cmd.GetType() {
	case CmdCounterInc:
		m.With(cmd.GetLabels()).Inc()
	case CmdCounterAdd:
		m.With(cmd.GetLabels()).Add(cmd.GetValue())
	default:
		s.logger.Error(
			"unable to find requested counter metric command",
			zap.Error(ErrMetricCmdNotFound),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
			zap.String(LoggerTagCmd, cmd.GetType().String()),
		)
	}
}

func (s *Service) handleGaugeCmd(cmd cmdForGauge) {
	m, err := s.GetGauge(cmd.GetMetricName())
	if err != nil {
		s.logger.Error(
			"unable to get the gauge metric",
			zap.Error(err),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
		)
		return
	}

	switch cmd.GetType() {
	case CmdGaugeSet:
		m.Set(cmd.GetValue())
	case CmdGaugeInc:
		m.Inc()
	case CmdGaugeDec:
		m.Dec()
	case CmdGaugeAdd:
		m.Add(cmd.GetValue())
	case CmdGaugeSub:
		m.Sub(cmd.GetValue())
	case CmdGaugeSetToCurrentTime:
		m.SetToCurrentTime()
	default:
		s.logger.Error(
			"unable to find requested gauge metric command",
			zap.Error(ErrMetricCmdNotFound),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
			zap.String(LoggerTagCmd, cmd.GetType().String()),
		)
	}
}

func (s *Service) handleGaugeVecCmd(cmd cmdForGaugeVec) {
	m, err := s.GetGaugeVec(cmd.GetMetricName())
	if err != nil {
		s.logger.Error(
			"unable to get the gauge vec metric",
			zap.Error(err),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
		)
		return
	}

	switch cmd.GetType() {
	case CmdGaugeSet:
		m.With(cmd.GetLabels()).Set(cmd.GetValue())
	case CmdGaugeInc:
		m.With(cmd.GetLabels()).Inc()
	case CmdGaugeDec:
		m.With(cmd.GetLabels()).Dec()
	case CmdGaugeAdd:
		m.With(cmd.GetLabels()).Add(cmd.GetValue())
	case CmdGaugeSub:
		m.With(cmd.GetLabels()).Sub(cmd.GetValue())
	case CmdGaugeSetToCurrentTime:
		m.With(cmd.GetLabels()).SetToCurrentTime()
	default:
		s.logger.Error(
			"unable to find requested gauge vec metric command",
			zap.Error(ErrMetricCmdNotFound),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
			zap.String(LoggerTagCmd, cmd.GetType().String()),
		)
	}
}

func (s *Service) handleObserverCmd(cmd cmdForObserver) {
	m, err := s.GetObserver(cmd.GetMetricName())
	if err != nil {
		s.logger.Error(
			"unable to get the observer metric",
			zap.Error(err),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
		)
		return
	}

	switch cmd.GetType() {
	case CmdObserverObserve:
		m.Observe(cmd.GetValue())
	default:
		s.logger.Error(
			"unable to find requested observer metric command",
			zap.Error(ErrMetricCmdNotFound),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
			zap.String(LoggerTagCmd, cmd.GetType().String()),
		)
	}
}

func (s *Service) handleObserverVecCmd(cmd cmdForObserverVec) {
	m, err := s.GetObserverVec(cmd.GetMetricName())
	if err != nil {
		s.logger.Error(
			"unable to get the observer vec metric",
			zap.Error(err),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
		)
		return
	}

	switch cmd.GetType() {
	case CmdObserverObserve:
		m.With(cmd.GetLabels()).Observe(cmd.GetValue())
	default:
		s.logger.Error(
			"unable to find requested observer vec metric command",
			zap.Error(ErrMetricCmdNotFound),
			zap.String(LoggerTagMetric, cmd.GetMetricName()),
			zap.String(LoggerTagCmd, cmd.GetType().String()),
		)
	}
}
