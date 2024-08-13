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
			Gauge:   0,
			Counter: 0,
		},
	}

	return repo
}

func (m *MemStorage) AddMetrics(newMetric models.MetricsData) {
	log.Printf("old value->%v", m.data)
	m.data.Gauge = newMetric.Gauge
	m.data.Counter += newMetric.Counter
	log.Printf("new value->%v", m.data)
	
}
