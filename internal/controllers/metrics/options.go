package metrics

import "strconv"

func (m *metrics) Add(code int, path string, timigs float64) {
	m.codes.WithLabelValues(strconv.Itoa(code), path).Inc()
	// m.timings.
}
