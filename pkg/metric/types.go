package metric

type config interface {
	GetAddress() string
	GetPath() string
}

type cmd interface {
	GetMetricName() string
}

type cmdForCounter interface {
	cmd
	GetType() CmdTypeCounter
	GetValue() float64
}

type cmdForCounterVec interface {
	cmdForCounter
	GetLabels() map[string]string
}

type cmdForGauge interface {
	cmd
	GetType() CmdTypeGauge
	GetValue() float64
}

type cmdForGaugeVec interface {
	cmdForGauge
	GetLabels() map[string]string
}

type cmdForObserver interface {
	cmd
	GetType() CmdTypeObserver
	GetValue() float64
}

type cmdForObserverVec interface {
	cmdForObserver
	GetLabels() map[string]string
}

type service interface {
	Init() error
	AddCounter(counterName, help string)
	AddCounterVec(counterName, help string, labelNames []string)
	AddCounterFunc(counterName, help string, callback func() float64)
	AddGauge(gaugeName, help string)
	AddGaugeVec(gaugeName, help string, labelNames []string)
	AddGaugeFunc(gaugeName, help string, callback func() float64)
	AddHistogram(histogramName, help string, buckets []float64)
	AddHistogramVec(histogramName, help string, labelNames []string, buckets []float64)
	AddSummary(summaryName, help string)
	AddSummaryVec(summaryName, help string, labelNames []string)

	CounterInc(counterName string)
	CounterAdd(counterName string, value float64)
	CounterVecInc(counterName string, labels map[string]string)
	CounterVecAdd(counterName string, labels map[string]string, value float64)
	GaugeSet(gaugeName string, value float64)
	GaugeInc(gaugeName string)
	GaugeDec(gaugeName string)
	GaugeAdd(gaugeName string, value float64)
	GaugeSub(gaugeName string, value float64)
	GaugeSetToCurrentTime(gaugeName string)
	GaugeVecSet(gaugeName string, labels map[string]string, value float64)
	GaugeVecInc(gaugeName string, labels map[string]string)
	GaugeVecDec(gaugeName string, labels map[string]string)
	GaugeVecAdd(gaugeName string, labels map[string]string, value float64)
	GaugeVecSub(gaugeName string, labels map[string]string, value float64)
	GaugeVecSetToCurrentTime(gaugeName string, labels map[string]string)
	ObserverObserve(observerName string, value float64)
	ObserverVecObserve(observerName string, labels map[string]string, value float64)
}
