// go:build unit
package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	repository "github.com/chinathaip/lmwn-tft-2024-assignment/data/repository_impl"
	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/controller"
	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/dto"
	"github.com/stretchr/testify/assert"
)

func TestGinRouter(t *testing.T) {
	covidData := &dto.CovidData{
		Data: []dto.CovidCase{
			{
				Province: "Bangkok",
				Age:      25,
			},
		},
	}
	expectedResponse := &dto.CovidSummary{
		Province: map[string]int{
			"Bangkok": 1,
		},
		AgeGroup: map[string]int{
			"0-30":  1,
			"31-60": 0,
			"60+":   0,
			"N/A":   0,
		},
	}
	covidRepository := repository.NewCovidRepositoryImpl(covidData)
	covidController := controller.NewCovidController(covidRepository)
	router := NewGinRouter(covidController)
	expectedCovidData, _ := json.Marshal(expectedResponse)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/covid/summary", nil)
	router.ServeHTTP(rec, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(expectedCovidData), rec.Body.String())
}
