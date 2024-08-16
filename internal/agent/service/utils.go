package service

import (
	"github.com/1Lindo/metrics_collector/internal/agent/models"
	"math/rand"
	"runtime"
)

func toMetricsData(rawData runtime.MemStats, syncCount int) models.MetricsData {
	return models.MetricsData{
		Gauge: map[string]float64{
			"Alloc":         float64(rawData.Alloc),
			"BuckHashSys":   float64(rawData.BuckHashSys),
			"Frees":         float64(rawData.Frees),
			"GCCPUFraction": float64(rawData.GCCPUFraction),
			"GCSys":         float64(rawData.GCSys),
			"HeapAlloc":     float64(rawData.HeapAlloc),
			"HeapIdle":      float64(rawData.HeapIdle),
			"HeapInuse":     float64(rawData.HeapInuse),
			"HeapObjects":   float64(rawData.HeapObjects),
			"HeapReleased":  float64(rawData.HeapReleased),
			"HeapSys":       float64(rawData.HeapSys),
			"LastGC":        float64(rawData.LastGC),
			"Lookups":       float64(rawData.Lookups),
			"MCacheInuse":   float64(rawData.MCacheInuse),
			"MCacheSys":     float64(rawData.MCacheSys),
			"MSpanInuse":    float64(rawData.MSpanInuse),
			"MSpanSys":      float64(rawData.MSpanSys),
			"Mallocs":       float64(rawData.Mallocs),
			"NextGC":        float64(rawData.NextGC),
			"NumForcedGC":   float64(rawData.NumForcedGC),
			"NumGC":         float64(rawData.NumGC),
			"OtherSys":      float64(rawData.OtherSys),
			"PauseTotalNs":  float64(rawData.PauseTotalNs),
			"StackInuse":    float64(rawData.StackInuse),
			"StackSys":      float64(rawData.StackSys),
			"Sys":           float64(rawData.Sys),
			"TotalAlloc":    float64(rawData.TotalAlloc),
			"RandomValue":   rand.Float64() * 1000,
		},
		Counter: map[string]int64{
			"PollCount": int64(syncCount),
		},
	}
}
