package repository

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
)

type MemStorage interface {
	AddMetrics(newMetric models.MetricsData, metricType string) bool
	GetAllMetrics() models.MetricsData
}

type memStorage struct {
	data models.MetricsData
}

// Симуляция хранилища метрик
func InitRepository() MemStorage {
	repo := &memStorage{
		data: models.MetricsData{
			Gauge:   map[string]float64{},
			Counter: map[string]int64{},
		},
	}

	return repo
}

func (m *memStorage) AddMetrics(newMetric models.MetricsData, metricType string) bool {

	if metricType == models.Gauge {
		for k, v := range newMetric.Gauge {
			m.data.Gauge[k] = v
		}
		return true
	}

	if metricType == models.Counter {
		for k, v := range newMetric.Counter {
			m.data.Counter[k] += v
		}
		return true
	}
	return false
}

func (m *memStorage) GetAllMetrics() models.MetricsData {
	return m.data
}
