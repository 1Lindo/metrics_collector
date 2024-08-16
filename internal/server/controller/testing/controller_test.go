package testing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/1Lindo/metrics_collector/internal/server/controller"
	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock service
type MockService struct {
	mock.Mock
}

func (m *MockService) AddMetrics(metric models.MetricsData, metricType string) bool {
	args := m.Called(metric, metricType)
	return args.Bool(0)
}

func (m *MockService) GetAllMetrics() models.MetricsData {
	args := m.Called()
	return args.Get(0).(models.MetricsData)
}

func Test_UpdateMetrics(t *testing.T) {
	mockService := new(MockService)
	ctrl := controller.InitController(mockService)

	tests := []struct {
		name           string
		method         string
		url            string
		contentType    string
		body           string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:           "Invalid HTTP method",
			method:         http.MethodGet,
			url:            "/update/gauge/test/10",
			contentType:    "text/plain",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Unsupported media type",
			method:         http.MethodPost,
			url:            "/update/gauge/test/10",
			contentType:    "application/json",
			expectedStatus: http.StatusUnsupportedMediaType,
		},
		{
			name:           "Invalid URL format",
			method:         http.MethodPost,
			url:            "/update/gauge/test/",
			contentType:    "text/plain",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:        "Successful gauge update",
			method:      http.MethodPost,
			url:         "/update/gauge/test/10",
			contentType: "text/plain",
			mockSetup: func() {
				mockService.On("AddMetrics", models.MetricsData{
					Gauge: map[string]float64{"test": 10.0},
				}, models.Gauge).Return(true)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Gauge conversion error",
			method:         http.MethodPost,
			url:            "/update/gauge/test/invalid",
			contentType:    "text/plain",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Unsupported metric type",
			method:         http.MethodPost,
			url:            "/update/unknown/test/10",
			contentType:    "text/plain",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
				return
			}

			req := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)

			rec := httptest.NewRecorder()

			ctrl.UpdateMetrics(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func Test_GetMetrics(t *testing.T) {
	mockService := new(MockService)
	ctrl := controller.InitController(mockService)

	expectedMetrics := models.MetricsData{
		Gauge:   map[string]float64{"test_gauge": 1.23},
		Counter: map[string]int64{"test_counter": 100},
	}

	mockService.On("GetAllMetrics").Return(expectedMetrics)

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()

	ctrl.GetMetrics(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var responseData models.MetricsData
	err := json.NewDecoder(res.Body).Decode(&responseData)
	assert.NoError(t, err)
	assert.Equal(t, expectedMetrics, responseData)
}
