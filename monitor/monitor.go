package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

var WebRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "web_request_total",
		Help: "Number of request in total",
	},
	// 设置标签
	[]string{"api"},
)

// web_request_duration_seconds，Histogram类型指标，bucket代表duration的分布区间
var WebRequestHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "web_request_histogram_seconds",
		Help:    "web request duration distribution",
		Buckets: []float64{10, 50, 100, 200, 500, 1000},
	},
	[]string{"api"},
)

var WebRequestSummery = prometheus.NewSummaryVec(prometheus.SummaryOpts{
	Name:       "web_request_duration_seconds",
	Help:       "web request duration seconds",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
},
	[]string{"api"},
)

func init() {
	// 注册监控指标
	prometheus.MustRegister(WebRequestTotal)
	prometheus.MustRegister(WebRequestHistogram)
	prometheus.MustRegister(WebRequestSummery)
}

func Monitor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		cost := time.Since(start)
		WebRequestTotal.With(prometheus.Labels{"api": r.URL.Path}).Inc()
		WebRequestHistogram.With(prometheus.Labels{"api": r.URL.Path}).Observe(float64(cost.Nanoseconds() / 1e6))
		WebRequestSummery.With(prometheus.Labels{"api": r.URL.Path}).Observe(float64(cost.Nanoseconds() / 1e6))
	}
}
