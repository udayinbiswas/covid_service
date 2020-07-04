package external

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://api.covid19india.org/data.json"
)

// CrowdSource ...
type CrowdSource struct {
	CasesTimeSeries []CasesTimeSeries `json:"cases_time_series"`
}

// CasesTimeSeries ...
type CasesTimeSeries struct {
	DailyConfirmed string `json:"dailyconfirmed"`
	DailyDeceased  string `json:"dailydeceased"`
	DailyRecovered string `json:"dailyrecovered"`
	Date           string `json:"date"`
	TotalConfirmed string `json:"totalconfirmed"`
	TotalDeceased  string `json:"totaldeceased"`
	TotalRecovered string `json:"totalrecovered"`
}

// GetCrowdSourcedData ...
func GetCrowdSourcedData() (*CrowdSource, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var cs CrowdSource
	err = json.Unmarshal(body, &cs)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
