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
	StateWise       []StateWise       `json:"statewise"`
	Tested          []Tested          `json:"tested"`
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

// StateWise ...
type StateWise struct {
	Active          string `json:"active"`
	Confirmed       string `json:"confirmed"`
	Deaths          string `json:"deaths"`
	DeltaConfirmed  string `json:"deltaconfirmed"`
	DeltaDeaths     string `json:"deltadeaths"`
	DeltaRecovered  string `json:"deltarecovered"`
	LastUpdatedTime string `json:"lastupdatetime"`
	MigratedOther   string `json:"migratedother"`
	Recovered       string `json:"recovered"`
	State           string `json:"state"`
	StateCode       string `json:"statecode"`
	StateNotes      string `json:"statenotes"`
}

// Tested ...
type Tested struct {
	IndividualsTestedPerConfirmedCase string `json:"individualstestedperconfirmedcase"`
	PositiveCasesFromSamplesReported  string `json:"positivecasesfromsamplesreported"`
	SampleReportedToday               string `json:"samplereportedtoday"`
	Source                            string `json:"source"`
	Source1                           string `json:"source1"`
	TestedAsOf                        string `json:"testedasof"`
	TestPositivityRate                string `json:"testpositivityrate"`
	TestsConductedByPrivateLabs       string `json:"testsconductedbyprivatelabs"`
	TestsPerConfirmedCase             string `json:"testsperconfirmedcase"`
	TestsPerMillion                   string `json:"testspermillion"`
	TotalIndividualsTested            string `json:"totalindividualstested"`
	TotalPositiveCases                string `json:"totalpositivecases"`
	TotalSamplesTested                string `json:"totalsamplestested"`
	UpdateTimestamp                   string `json:"updatetimestamp"`
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
