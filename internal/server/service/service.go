package service

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/repository"
)

type Service struct {
	repo *repository.MemStorage
}

func InitCollectorService(repo *repository.MemStorage) *Service {
	srv := &Service{
		repo: repo,
	}

	return srv
}

func (s *Service) AddMetrics(metric models.MetricsData) {
	s.repo.AddMetrics(metric)
}
