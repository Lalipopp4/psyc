package metrics

import "strconv"

func (m *metrics) Add(code int, path string, timing float64) {
	m.codes.WithLabelValues(strconv.Itoa(code), path).Inc()
	m.timings.WithLabelValues(path).Observe(timing)
}
