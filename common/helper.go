package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/dto"
)

func FetchCovidCase(dataSource string) (*dto.CovidData, error) {
	// ensure covid-cases.json is present
	if _, err := os.Open("covid-cases.json"); err != nil {
		if !os.IsNotExist(err) {
			return nil, NewFetchCovidCaseError(fmt.Sprintf("error while opening covid-cases.json: %v", err))
		}

		resp, err := http.Get(dataSource)
		if err != nil {
			return nil, NewFetchCovidCaseError(fmt.Sprintf("error while fetching covid data source: %v", err))
		}

		respBody, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, NewFetchCovidCaseError(fmt.Sprintf("error while parsing response body: %v", err))
		}

		if err = os.WriteFile("covid-cases.json", respBody, 0755); err != nil {
			return nil, NewFetchCovidCaseError(fmt.Sprintf("error while writing covid case to file: %v", err))
		}
	}

	// parse covid-cases.json
	fileContent, err := os.ReadFile("covid-cases.json")
	if err != nil {
		return nil, NewFetchCovidCaseError(fmt.Sprintf("error while reading covid-cases.json: %v", err))
	}

	var covidData dto.CovidData
	if err := json.Unmarshal(fileContent, &covidData); err != nil {
		return nil, NewFetchCovidCaseError(fmt.Sprintf("error while unmarshalling covid data: %v", err))
	}

	return &covidData, nil
}
