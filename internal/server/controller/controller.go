package controller

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/service"
	"net/http"
	"strconv"
	"strings"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
	Empty   = ""
)

type Controller struct {
	srv *service.Service
}

func InitController(s *service.Service) Controller {
	return Controller{
		srv: s,
	}
}

func (c Controller) UpdateMetrics(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if req.Method != http.MethodPost {
		http.Error(res, "Status Method Is Not Allowed", http.StatusMethodNotAllowed)
	}
	if contentType != "text/plain" {
		http.Error(res, "Unsupported Media Type", http.StatusUnsupportedMediaType)
	}

	parts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(parts) != 4 {
		http.Error(res, "Invalid URL format", http.StatusNotFound)
	}

	metricType, metricName, metricValue := parts[1], parts[2], parts[3]

	// Проверка имени метрики
	if metricName == Empty {
		http.Error(res, "Metric name is required", http.StatusNotFound)
	}

	var err error
	switch metricType {
	case Gauge:
		gauge, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(res, "Ошибка конвертации", http.StatusInternalServerError)
		}
		metric := models.MetricsData{
			Gauge: gauge,
		}
		c.srv.AddMetrics(metric)
		res.WriteHeader(http.StatusOK)
	case Counter:
		counter, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "Ошибка конвертации", http.StatusInternalServerError)
		}
		metric := models.MetricsData{
			Counter: counter,
		}
		c.srv.AddMetrics(metric)
		res.WriteHeader(http.StatusOK)
	default:
		http.Error(res, "Неподдерживаемый тип метрики", http.StatusBadRequest)
	}
	if err != nil {
		http.Error(res, "Invalid value", http.StatusBadRequest)
	}
}
