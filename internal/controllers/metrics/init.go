package metrics

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	codes   *prometheus.CounterVec
	timings *prometheus.HistogramVec
}

func New(codes *prometheus.CounterVec, timings *prometheus.HistogramVec) Metrics {
	return &metrics{codes, timings}
}
