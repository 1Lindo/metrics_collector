package repository

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
)

type MemStorage struct {
	data models.MetricsData
}

// Симуляция хранилища метрик
func InitRepository() *MemStorage {
	repo := &MemStorage{
		data: models.MetricsData{
			Gauge:   map[string]float64{},
			Counter: map[string]int64{},
		},
	}

	return repo
}

func (m *MemStorage) AddMetrics(newMetric models.MetricsData, metricType string) {

	if metricType == models.Gauge {
		for k, v := range newMetric.Gauge {
			m.data.Gauge[k] = v
		}
	}

	if metricType == models.Counter {
		for k, v := range newMetric.Counter {
			m.data.Counter[k] += v
		}
	}
}

func (m *MemStorage) GetAllMetrics() models.MetricsData {
	return m.data
}
