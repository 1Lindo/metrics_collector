package repository

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"log"
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
	log.Printf("old value -> %v", m.data)

	if metricType == models.Gauge {
		for k, v := range newMetric.Gauge {
			m.data.Gauge[k] = v
		}
	}

	for k, v := range newMetric.Counter {
		if _, ok := m.data.Counter[k]; ok {
			m.data.Counter[k] += v
		} else {
			m.data.Counter[k] = v
		}
	}

	log.Printf("new value -> %v", m.data)
}
