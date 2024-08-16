package models

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type MetricsData struct {
	Gauge   map[string]float64
	Counter map[string]int64
}
