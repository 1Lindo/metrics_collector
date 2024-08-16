package testing

import (
	"testing"

	"github.com/1Lindo/metrics_collector/internal/server/models"
	"github.com/1Lindo/metrics_collector/internal/server/repository"
	"github.com/stretchr/testify/assert"
)

func TestAddMetrics(t *testing.T) {
	tests := []struct {
		name            string
		metricType      string
		inputMetric     models.MetricsData
		expectedGauge   map[string]float64
		expectedCounter map[string]int64
		expectedResult  bool
	}{
		{
			name:       "Add gauge metric",
			metricType: models.Gauge,
			inputMetric: models.MetricsData{
				Gauge: map[string]float64{"metric1": 10.5},
			},
			expectedGauge:   map[string]float64{"metric1": 10.5},
			expectedCounter: map[string]int64{},
			expectedResult:  true,
		},
		{
			name:       "Add counter metric",
			metricType: models.Counter,
			inputMetric: models.MetricsData{
				Counter: map[string]int64{"metric2": 3},
			},
			expectedGauge:   map[string]float64{},
			expectedCounter: map[string]int64{"metric2": 3},
			expectedResult:  true,
		},
		{
			name:       "Add to existing counter metric",
			metricType: models.Counter,
			inputMetric: models.MetricsData{
				Counter: map[string]int64{"metric2": 2},
			},
			expectedGauge:   map[string]float64{},
			expectedCounter: map[string]int64{"metric2": 2},
			expectedResult:  true,
		},
		{
			name:       "Invalid metric type",
			metricType: "unknown",
			inputMetric: models.MetricsData{
				Gauge: map[string]float64{"metric3": 1.1},
			},
			expectedGauge:   map[string]float64{},
			expectedCounter: map[string]int64{},
			expectedResult:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.InitRepository()

			result := repo.AddMetrics(tt.inputMetric, tt.metricType)
			assert.Equal(t, tt.expectedResult, result)

			data := repo.GetAllMetrics()
			assert.Equal(t, tt.expectedGauge, data.Gauge)
			assert.Equal(t, tt.expectedCounter, data.Counter)
		})
	}
}

func TestGetAllMetrics(t *testing.T) {
	repo := repository.InitRepository()

	// Add some metrics
	repo.AddMetrics(models.MetricsData{
		Gauge: map[string]float64{"metric1": 10.5},
	}, models.Gauge)

	repo.AddMetrics(models.MetricsData{
		Counter: map[string]int64{"metric2": 5},
	}, models.Counter)

	expectedMetrics := models.MetricsData{
		Gauge:   map[string]float64{"metric1": 10.5},
		Counter: map[string]int64{"metric2": 5},
	}

	result := repo.GetAllMetrics()

	assert.Equal(t, expectedMetrics, result)
}
