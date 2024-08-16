package controller

import (
	"encoding/json"
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/service"
	"net/http"
	"strconv"
	"strings"
)

type Controller interface {
	UpdateMetrics(res http.ResponseWriter, req *http.Request)
	GetMetrics(res http.ResponseWriter, req *http.Request)
}
type controller struct {
	srv service.Service
}

func InitController(s service.Service) Controller {
	return controller{
		srv: s,
	}
}

func (c controller) UpdateMetrics(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if req.Method != http.MethodPost {
		http.Error(res, "Status Method Is Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if contentType != "text/plain" {
		http.Error(res, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
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
			return
		}
		metric := models.MetricsData{
			Gauge: map[string]float64{metricName: gauge},
		}
		if ok := c.srv.AddMetrics(metric, metricType); !ok {
			http.Error(res, "Repository err", http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusOK)
	case models.Counter:
		counter, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "Converting error, check metric value type", http.StatusBadRequest)
			return
		}
		metric := models.MetricsData{
			Counter: map[string]int64{metricName: counter},
		}
		if ok := c.srv.AddMetrics(metric, metricType); !ok {
			http.Error(res, "Repository err", http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusOK)
	default:
		http.Error(res, "Unsupported metric type", http.StatusBadRequest)
		return
	}
}

func (c controller) GetMetrics(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Status Method Is Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data := c.srv.GetAllMetrics()
	dataJSON, err := json.Marshal(&data)
	if err != nil {
		http.Error(res, "JSON parsing error", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	if _, err := res.Write(dataJSON); err != nil {
		http.Error(res, "Sending response error", http.StatusInternalServerError)
		return
	}
}
