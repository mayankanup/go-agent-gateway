package metrics

import "sync/atomic"

type Metrics struct {
	RequestCount uint64
	ToolCalls    uint64
}

func (m *Metrics) IncrementRequests() {
	atomic.AddUint64(
		&m.RequestCount,
		1,
	)
}

func (m *Metrics) IncrementToolCalls() {
	atomic.AddUint64(
		&m.ToolCalls,
		1,
	)
}
