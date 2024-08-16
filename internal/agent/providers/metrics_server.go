package providers

import (
	"bytes"
	"fmt"
	"github.com/1Lindo/metrics_collector/internal/agent/models"
	"log"
	"net/http"
)

type MetricsProvider struct {
}

func InitMetricsProvider() MetricsProvider {
	return MetricsProvider{}
}

func (m *MetricsProvider) SendValue(data models.MetricsData) {
	url := "http://localhost:8080/update"

	for k, v := range data.Gauge {
		newURL := fmt.Sprintf(url+"/%v/%v/%v", models.Gauge, k, v)
		resp, err := http.Post(newURL, "text/plain", bytes.NewBuffer([]byte("")))
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
		newURLCounter := fmt.Sprintf(url+"/%v/%v/%v", models.Counter, k, v)
		resp, err := http.Post(newURLCounter, "text/plain", bytes.NewBuffer([]byte("")))
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
