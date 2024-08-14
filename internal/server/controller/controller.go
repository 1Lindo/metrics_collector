package controller

import (
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/service"
	"net/http"
	"strconv"
	"strings"
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
		return
	}

	metricType, metricName, metricValue := parts[1], parts[2], parts[3]

	if metricName == models.Empty {
		http.Error(res, "Metric name is required", http.StatusNotFound)
		return
	}

	switch metricType {
	case models.Gauge:
		gauge, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(res, "Converting error, check metric value type", http.StatusBadRequest)
		}
		metric := models.MetricsData{
			Gauge: map[string]float64{metricName: gauge},
		}
		c.srv.AddMetrics(metric, metricType)
		res.WriteHeader(http.StatusOK)
	case models.Counter:
		counter, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "Converting error, check metric value type", http.StatusBadRequest)
		}
		metric := models.MetricsData{
			Counter: map[string]int64{metricName: counter},
		}
		c.srv.AddMetrics(metric, metricType)
		res.WriteHeader(http.StatusOK)
	default:
		http.Error(res, "Unsupported metric type", http.StatusBadRequest)
	}
}
