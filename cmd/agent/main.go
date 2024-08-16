package main

import (
	"bytes"
	"fmt"
	"github.com/1Lindo/metrics_collector/internal/agent/models"
	"github.com/1Lindo/metrics_collector/internal/agent/service"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	srv := service.InitSrv()
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
			sendValue(data)
			count = 0
			mu.Unlock()
		}
	}

}

func sendValue(data models.MetricsData) {
	url := "http://localhost:8080/update"

	for k, v := range data.Gauge {
		newUrl := fmt.Sprintf(url+"/%v/%v/%v", models.Gauge, k, v)
		resp, err := http.Post(newUrl, "text/plain", bytes.NewBuffer([]byte("")))
		if err != nil {
			log.Printf("Ошибка при отправке [Gauge:%v | Value:%v], err:%v", k, v, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Printf("Значение [Gauge:%v | Value:%v] успешно отправлено:", k, v)
		} else {
			log.Printf("Ошибка при отправке [Gauge:%v | Value:%v], статус:%v", k, v, resp.Status)
		}
	}
	for k, v := range data.Counter {
		newUrlCounter := fmt.Sprintf(url+"/%v/%v/%v", models.Counter, k, v)
		resp, err := http.Post(newUrlCounter, "text/plain", bytes.NewBuffer([]byte("")))
		if err != nil {
			log.Printf("Ошибка при отправке [Counter:%v | Value:%v], err:%v", k, v, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Printf("Значение [Counter:%v | Value:%v] успешно отправлено:", k, v)
		} else {
			log.Printf("Ошибка при отправке [Counter:%v | Value:%v], статус:%v", k, v, resp.Status)
		}
	}
}
