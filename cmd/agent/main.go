package main

import (
	"github.com/1Lindo/metrics_collector/internal/agent/models"
	"github.com/1Lindo/metrics_collector/internal/agent/providers"
	"github.com/1Lindo/metrics_collector/internal/agent/service"
	"sync"
	"time"
)

func main() {
	srv := service.InitSrv()
	prv := providers.InitMetricsProvider()

	count := 0

	var mu sync.Mutex
	data := models.MetricsData{}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			mu.Lock()
			count++
			data = srv.CollectMetrics(count)
			mu.Unlock()
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case <-ticker.C:
			mu.Lock()
			prv.SendValue(data)
			count = 0
			mu.Unlock()
		}
	}

}
