package restapi

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//OpsProcessed global for metrics
var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

//RecordMetrics start the timer
func RecordMetrics() {
	go func() {
		for {
			OpsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
