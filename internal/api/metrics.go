package api

import (
	"encoding/json"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/metrics"
)

type MetricsHandler struct {
	metrics *metrics.Metrics
}

func NewMetricsHandler(
	m *metrics.Metrics,
) *MetricsHandler {

	return &MetricsHandler{
		metrics: m,
	}
}

func (h *MetricsHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).Encode(
		map[string]any{
			"requests": h.metrics.RequestCount,

			"toolCalls": h.metrics.ToolCalls,
		},
	)
}
