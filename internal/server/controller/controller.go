package controller

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/service"
	"net/http"
	"strconv"
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
	metricTypeStr := req.URL.Query().Get("type")
	metricNameStr := req.URL.Query().Get("name")
	metricValueStr := req.URL.Query().Get("value")

	switch metricNameStr {
	case Gauge:
		if metricTypeStr != "float64" {
			http.Error(res, "Ошибка типа данных", http.StatusInternalServerError)
		}
		gauge, err := strconv.ParseFloat(metricValueStr, 64)
		if err != nil {
			http.Error(res, "Ошибка конвертации", http.StatusInternalServerError)
		}
		metric := models.MetricsData{
			Gauge: gauge,
		}
		c.srv.AddMetrics(metric)
		res.WriteHeader(http.StatusOK)
	case Counter:
		if metricTypeStr != "int64" {
			http.Error(res, "Ошибка типа данных", http.StatusInternalServerError)
		}
		counter, err := strconv.ParseInt(metricValueStr, 10, 64)
		if err != nil {
			http.Error(res, "Ошибка конвертации", http.StatusInternalServerError)
		}
		metric := models.MetricsData{
			Counter: counter,
		}
		c.srv.AddMetrics(metric)
		res.WriteHeader(http.StatusOK)
	case Empty:
		http.Error(res, "Отсутствует название метрики", http.StatusNotFound)
	default:
		http.Error(res, "Неподдерживаемая метрика", http.StatusBadRequest)
	}
}
