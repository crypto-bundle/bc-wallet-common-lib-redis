package metric

func (s *Service) CounterInc(counterName string) {
	go func() {
		s.cmdToCountersCh <- &CmdForCounter{metricName: s.formatMetricName(counterName), cmdType: CmdCounterInc}
	}()
}

func (s *Service) CounterAdd(counterName string, value float64) {
	go func() {
		s.cmdToCountersCh <- &CmdForCounter{metricName: s.formatMetricName(counterName), cmdType: CmdCounterAdd, addValue: value}
	}()
}

func (s *Service) CounterVecInc(counterName string, labels map[string]string) {
	go func() {
		s.cmdToCounterVecsCh <- &CmdForCounterVec{
			labels: labels,
			CmdForCounter: &CmdForCounter{
				metricName: s.formatMetricName(counterName),
				cmdType:    CmdCounterInc,
			},
		}
	}()
}

func (s *Service) CounterVecAdd(counterName string, labels map[string]string, value float64) {
	go func() {
		s.cmdToCounterVecsCh <- &CmdForCounterVec{
			labels: labels,
			CmdForCounter: &CmdForCounter{
				metricName: s.formatMetricName(counterName),
				cmdType:    CmdCounterAdd,
				addValue:   value,
			},
		}
	}()
}

func (s *Service) GaugeSet(gaugeName string, value float64) {
	go func() {
		s.cmdToGaugesCh <- &CmdForGauge{metricName: s.formatMetricName(gaugeName), cmdType: CmdGaugeSet, value: value}
	}()
}

func (s *Service) GaugeInc(gaugeName string) {
	go func() {
		s.cmdToGaugesCh <- &CmdForGauge{metricName: s.formatMetricName(gaugeName), cmdType: CmdGaugeInc}
	}()
}

func (s *Service) GaugeDec(gaugeName string) {
	go func() {
		s.cmdToGaugesCh <- &CmdForGauge{metricName: s.formatMetricName(gaugeName), cmdType: CmdGaugeDec}
	}()
}

func (s *Service) GaugeAdd(gaugeName string, value float64) {
	go func() {
		s.cmdToGaugesCh <- &CmdForGauge{metricName: s.formatMetricName(gaugeName), cmdType: CmdGaugeAdd, value: value}
	}()
}

func (s *Service) GaugeSub(gaugeName string, value float64) {
	go func() {
		s.cmdToGaugesCh <- &CmdForGauge{metricName: s.formatMetricName(gaugeName), cmdType: CmdGaugeSub, value: value}
	}()
}

func (s *Service) GaugeSetToCurrentTime(gaugeName string) {
	go func() {
		s.cmdToGaugesCh <- &CmdForGauge{metricName: s.formatMetricName(gaugeName), cmdType: CmdGaugeSetToCurrentTime}
	}()
}

func (s *Service) GaugeVecSet(gaugeName string, labels map[string]string, value float64) {
	go func() {
		s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
			labels: labels,
			CmdForGauge: &CmdForGauge{
				metricName: s.formatMetricName(gaugeName),
				cmdType:    CmdGaugeSet,
				value:      value,
			},
		}
	}()
}

func (s *Service) GaugeVecInc(gaugeName string, labels map[string]string) {
	go func() {
		s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
			labels: labels,
			CmdForGauge: &CmdForGauge{
				metricName: s.formatMetricName(gaugeName),
				cmdType:    CmdGaugeInc,
			},
		}
	}()
}

func (s *Service) GaugeVecDec(gaugeName string, labels map[string]string) {
	go func() {
		s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
			labels: labels,
			CmdForGauge: &CmdForGauge{
				metricName: s.formatMetricName(gaugeName),
				cmdType:    CmdGaugeDec,
			},
		}
	}()
}

func (s *Service) GaugeVecAdd(gaugeName string, labels map[string]string, value float64) {
	go func() {
		s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
			labels: labels,
			CmdForGauge: &CmdForGauge{
				metricName: s.formatMetricName(gaugeName),
				cmdType:    CmdGaugeAdd,
				value:      value,
			},
		}
	}()
}

func (s *Service) GaugeVecSub(gaugeName string, labels map[string]string, value float64) {
	go func() {
		s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
			labels: labels,
			CmdForGauge: &CmdForGauge{
				metricName: s.formatMetricName(gaugeName),
				cmdType:    CmdGaugeSub,
				value:      value,
			},
		}
	}()
}

func (s *Service) GaugeVecSetToCurrentTime(gaugeName string, labels map[string]string) {
	go func() {
		s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
			labels: labels,
			CmdForGauge: &CmdForGauge{
				metricName: s.formatMetricName(gaugeName),
				cmdType:    CmdGaugeSetToCurrentTime,
			},
		}
	}()
}

func (s *Service) ObserverObserve(observerName string, value float64) {
	go func() {
		s.cmdToObserversCh <- &CmdForObserver{
			metricName: s.formatMetricName(observerName),
			cmdType:    CmdObserverObserve,
			value:      value,
		}
	}()
}

func (s *Service) ObserverVecObserve(observerName string, labels map[string]string, value float64) {
	go func() {
		s.cmdToObserverVecsCh <- &CmdForObserverVec{
			labels: labels,
			CmdForObserver: &CmdForObserver{
				metricName: s.formatMetricName(observerName),
				cmdType:    CmdObserverObserve,
				value:      value,
			},
		}
	}()
}
