// go:build unit
package repository_impl

import (
	"testing"

	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/dto"
	"github.com/stretchr/testify/assert"
)

type covidTest struct {
	name          string
	covidData     *dto.CovidData
	expectedYoung int
	expectedAdult int
	expectedOld   int
	expectedNA    int
}

func TestSummarize(t *testing.T) {
	t.Run("Test Mapping Age Group", func(t *testing.T) {
		tests := []covidTest{{
			name: "properly handle age group",
			covidData: &dto.CovidData{
				Data: []dto.CovidCase{
					{
						Age: 21,
					},
					{
						Age: 25,
					},
					{
						Age: 55,
					},
					{
						Age: 45,
					},
					{
						Age: 45,
					},
					{
						Age: 90,
					},
					{
						Age: 119,
					},
				}},
			expectedYoung: 2,
			expectedAdult: 3,
			expectedOld:   2,
		},
			{
				name: "edge cases should be put in N/A",
				covidData: &dto.CovidData{
					Data: []dto.CovidCase{
						{
							Age: -1,
						},
						{
							Age: 0,
						},
						{
							Age: 30,
						},
						{
							Age: 31,
						},
						{
							Age: 60,
						},
						{
							Age: 61,
						},
						{
							Age: 120,
						},
					}},
				expectedYoung: 2,
				expectedAdult: 2,
				expectedOld:   1,
				expectedNA:    2,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				covidRepoImpl := NewCovidRepositoryImpl(test.covidData)
				covidSummary := covidRepoImpl.Summarize()

				assert.Equal(t, test.expectedYoung, covidSummary.AgeGroup["0-30"])
				assert.Equal(t, test.expectedAdult, covidSummary.AgeGroup["31-60"])
				assert.Equal(t, test.expectedOld, covidSummary.AgeGroup["60+"])
				assert.Equal(t, test.expectedNA, covidSummary.AgeGroup["N/A"])
			})
		}
	})
}
