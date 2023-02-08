package metric

const (
	CmdIncName              = "Inc"
	CmdAddName              = "Add"
	CmdSetName              = "Set"
	CmdDecName              = "Dec"
	CmdSubName              = "Sub"
	CmdSetToCurrentTimeName = "SetToCurrentTime"
	CmdObserveName          = "Observe"
)

const (
	CmdCounterInc CmdTypeCounter = iota + 1
	CmdCounterAdd
)

const (
	CmdGaugeSet CmdTypeGauge = iota + 1
	CmdGaugeInc
	CmdGaugeDec
	CmdGaugeAdd
	CmdGaugeSub
	CmdGaugeSetToCurrentTime
)

const (
	CmdObserverObserve CmdTypeObserver = iota + 1
)

type (
	CmdTypeCounter  int
	CmdTypeGauge    int
	CmdTypeObserver int
)

func (t CmdTypeCounter) String() string {
	return [...]string{"",
		CmdIncName,
		CmdAddName,
	}[t]
}

func (t CmdTypeGauge) String() string {
	return [...]string{"",
		CmdSetName,
		CmdIncName,
		CmdDecName,
		CmdAddName,
		CmdSubName,
		CmdSetToCurrentTimeName,
	}[t]
}

func (t CmdTypeObserver) String() string {
	return [...]string{"",
		CmdObserveName,
	}[t]
}

type CmdForCounter struct {
	metricName string
	cmdType    CmdTypeCounter
	addValue   float64
}

func (c *CmdForCounter) GetMetricName() string {
	return c.metricName
}

func (c *CmdForCounter) GetType() CmdTypeCounter {
	return c.cmdType
}

func (c *CmdForCounter) GetValue() float64 {
	return c.addValue
}

type CmdForCounterVec struct {
	labels map[string]string

	*CmdForCounter
}

func (c *CmdForCounterVec) GetLabels() map[string]string {
	return c.labels
}

type CmdForGauge struct {
	metricName string
	cmdType    CmdTypeGauge
	value      float64
}

func (c *CmdForGauge) GetMetricName() string {
	return c.metricName
}

func (c *CmdForGauge) GetType() CmdTypeGauge {
	return c.cmdType
}

func (c *CmdForGauge) GetValue() float64 {
	return c.value
}

type CmdForGaugeVec struct {
	labels map[string]string

	*CmdForGauge
}

func (c *CmdForGaugeVec) GetLabels() map[string]string {
	return c.labels
}

type CmdForObserver struct {
	metricName string
	cmdType    CmdTypeObserver
	value      float64
}

func (c *CmdForObserver) GetMetricName() string {
	return c.metricName
}

func (c *CmdForObserver) GetType() CmdTypeObserver {
	return c.cmdType
}

func (c *CmdForObserver) GetValue() float64 {
	return c.value
}

type CmdForObserverVec struct {
	labels map[string]string

	*CmdForObserver
}

func (c *CmdForObserverVec) GetLabels() map[string]string {
	return c.labels
}
