package repository_impl

import (
	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/dto"
)

type CovidRepositoryImpl struct {
	covidData *dto.CovidData
}

func NewCovidRepositoryImpl(covidData *dto.CovidData) *CovidRepositoryImpl {
	return &CovidRepositoryImpl{covidData: covidData}
}

func (cr *CovidRepositoryImpl) Summarize() dto.CovidSummary {
	summary := dto.CovidSummary{
		Province: make(map[string]int),
		AgeGroup: map[string]int{
			"0-30":  0,
			"31-60": 0,
			"60+":   0,
			"N/A":   0,
		},
	}

	for _, data := range cr.covidData.Data {
		if data.Province != "" {
			summary.Province[data.Province]++
		}

		if 0 <= data.Age && data.Age <= 30 {
			summary.AgeGroup["0-30"]++
			continue
		}

		if 31 <= data.Age && data.Age <= 60 {
			summary.AgeGroup["31-60"]++
			continue
		}

		if 60 < data.Age && data.Age < 120 {
			summary.AgeGroup["60+"]++
			continue
		}

		summary.AgeGroup["N/A"]++
	}

	return summary
}
