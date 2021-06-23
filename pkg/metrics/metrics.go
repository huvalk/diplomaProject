package metrics

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "Requests count",
	}, []string{"method", "path", "status"})

	RequestCurrent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "request_current",
		Help: "Number of current requests",
	}, []string{"method", "path"})

	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_duration",
		Help: "Requests duration in second",
	}, []string{"method", "path"})

	MemoryPercent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memory_percent",
		Help: "Memory percent",
	}, []string{"percent"})
)

func RegisterPrometheus(e *echo.Echo) {
	//e.GET("/metrics", echo.HandlerFunc(promhttp.Handler))

	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestCurrent)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(MemoryPercent)
}
