package metrics
 
type Metrics interface {
	Add(code int, path string, timing float64)
} 