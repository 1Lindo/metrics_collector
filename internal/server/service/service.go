package service

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/repository"
)

type Service interface {
	AddMetrics(metric models.MetricsData, metricType string) bool
	GetAllMetrics() models.MetricsData
}

type service struct {
	repo repository.MemStorage
}

func InitCollectorService(repo repository.MemStorage) Service {
	srv := service{
		repo: repo,
	}

	return srv
}

func (s service) AddMetrics(metric models.MetricsData, metricType string) bool {
	if ok := s.repo.AddMetrics(metric, metricType); !ok {
		return false
	}
	return true
}

func (s service) GetAllMetrics() models.MetricsData {
	return s.repo.GetAllMetrics()
}
