package service

import (
	"github.com/1Lindo/metrics_collector/internal/agent/models"
	"runtime"
)

type Service struct {
}

func InitSrv() *Service {
	return &Service{}
}

func (s *Service) CollectMetrics(syncCount int) models.MetricsData {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return toMetricsData(memStats, syncCount)
}
