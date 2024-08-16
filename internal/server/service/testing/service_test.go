package testing

import (
	"testing"

	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockMemStorage struct {
	mock.Mock
}

func (m *MockMemStorage) AddMetrics(newMetric models.MetricsData, metricType string) bool {
	args := m.Called(newMetric, metricType)
	return args.Bool(0)
}

func (m *MockMemStorage) GetAllMetrics() models.MetricsData {
	args := m.Called()
	return args.Get(0).(models.MetricsData)
}

func TestAddMetrics(t *testing.T) {
	mockRepo := new(MockMemStorage)
	srv := service.InitCollectorService(mockRepo)

	metric := models.MetricsData{
		Gauge: map[string]float64{"metric1": 10.5},
	}
	mockRepo.On("AddMetrics", metric, "gauge").Return(true)

	ok := srv.AddMetrics(metric, "gauge")
	assert.True(t, ok)

	mockRepo.AssertExpectations(t)
}

func TestAddMetricsFailure(t *testing.T) {
	mockRepo := new(MockMemStorage)
	srv := service.InitCollectorService(mockRepo)

	metric := models.MetricsData{
		Gauge: map[string]float64{"metric1": 10.5},
	}
	mockRepo.On("AddMetrics", metric, "gauge").Return(false)

	ok := srv.AddMetrics(metric, "gauge")
	assert.False(t, ok)

	mockRepo.AssertExpectations(t)
}

func TestGetAllMetrics(t *testing.T) {
	mockRepo := new(MockMemStorage)
	srv := service.InitCollectorService(mockRepo)

	expectedData := models.MetricsData{
		Gauge: map[string]float64{"metric1": 10.5},
	}
	mockRepo.On("GetAllMetrics").Return(expectedData)

	data := srv.GetAllMetrics()
	assert.Equal(t, expectedData, data)

	mockRepo.AssertExpectations(t)
}
