package metric

func (s *Service) CounterInc(counterName string) {
	s.cmdToCountersCh <- &CmdForCounter{metricName: counterName, cmdType: CmdCounterInc}
}

func (s *Service) CounterAdd(counterName string, value float64) {
	s.cmdToCountersCh <- &CmdForCounter{metricName: counterName, cmdType: CmdCounterAdd, addValue: value}
}

func (s *Service) CounterVecInc(counterName string, labels map[string]string) {
	s.cmdToCounterVecsCh <- &CmdForCounterVec{
		labels: labels,
		CmdForCounter: &CmdForCounter{
			metricName: counterName,
			cmdType:    CmdCounterInc,
		},
	}
}

func (s *Service) CounterVecAdd(counterName string, labels map[string]string, value float64) {
	s.cmdToCounterVecsCh <- &CmdForCounterVec{
		labels: labels,
		CmdForCounter: &CmdForCounter{
			metricName: counterName,
			cmdType:    CmdCounterAdd,
			addValue:   value,
		},
	}
}

func (s *Service) GaugeSet(gaugeName string, value float64) {
	s.cmdToGaugesCh <- &CmdForGauge{metricName: gaugeName, cmdType: CmdGaugeSet, value: value}
}

func (s *Service) GaugeInc(gaugeName string) {
	s.cmdToGaugesCh <- &CmdForGauge{metricName: gaugeName, cmdType: CmdGaugeInc}
}

func (s *Service) GaugeDec(gaugeName string) {
	s.cmdToGaugesCh <- &CmdForGauge{metricName: gaugeName, cmdType: CmdGaugeDec}
}

func (s *Service) GaugeAdd(gaugeName string, value float64) {
	s.cmdToGaugesCh <- &CmdForGauge{metricName: gaugeName, cmdType: CmdGaugeAdd, value: value}
}

func (s *Service) GaugeSub(gaugeName string, value float64) {
	s.cmdToGaugesCh <- &CmdForGauge{metricName: gaugeName, cmdType: CmdGaugeSub, value: value}
}

func (s *Service) GaugeSetToCurrentTime(gaugeName string) {
	s.cmdToGaugesCh <- &CmdForGauge{metricName: gaugeName, cmdType: CmdGaugeSetToCurrentTime}
}

func (s *Service) GaugeVecSet(gaugeName string, labels map[string]string, value float64) {
	s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
		labels: labels,
		CmdForGauge: &CmdForGauge{
			metricName: gaugeName,
			cmdType:    CmdGaugeSet,
			value:      value,
		},
	}
}

func (s *Service) GaugeVecInc(gaugeName string, labels map[string]string) {
	s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
		labels: labels,
		CmdForGauge: &CmdForGauge{
			metricName: gaugeName,
			cmdType:    CmdGaugeInc,
		},
	}
}

func (s *Service) GaugeVecDec(gaugeName string, labels map[string]string) {
	s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
		labels: labels,
		CmdForGauge: &CmdForGauge{
			metricName: gaugeName,
			cmdType:    CmdGaugeDec,
		},
	}
}

func (s *Service) GaugeVecAdd(gaugeName string, labels map[string]string, value float64) {
	s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
		labels: labels,
		CmdForGauge: &CmdForGauge{
			metricName: gaugeName,
			cmdType:    CmdGaugeAdd,
			value:      value,
		},
	}
}

func (s *Service) GaugeVecSub(gaugeName string, labels map[string]string, value float64) {
	s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
		labels: labels,
		CmdForGauge: &CmdForGauge{
			metricName: gaugeName,
			cmdType:    CmdGaugeSub,
			value:      value,
		},
	}
}

func (s *Service) GaugeVecSetToCurrentTime(gaugeName string, labels map[string]string) {
	s.cmdToGaugeVecsCh <- &CmdForGaugeVec{
		labels: labels,
		CmdForGauge: &CmdForGauge{
			metricName: s.formatMetricName(gaugeName),
			cmdType:    CmdGaugeSetToCurrentTime,
		},
	}
}

func (s *Service) ObserverObserve(observerName string, value float64) {
	s.cmdToObserversCh <- &CmdForObserver{
		metricName: s.formatMetricName(observerName),
		cmdType:    CmdObserverObserve,
		value:      value,
	}
}

func (s *Service) ObserverVecObserve(observerName string, labels map[string]string, value float64) {
	s.cmdToObserverVecsCh <- &CmdForObserverVec{
		labels: labels,
		CmdForObserver: &CmdForObserver{
			metricName: s.formatMetricName(observerName),
			cmdType:    CmdObserverObserve,
			value:      value,
		},
	}
}
