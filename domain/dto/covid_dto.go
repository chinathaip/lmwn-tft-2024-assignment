package dto

type CovidData struct {
	Data []CovidCase `json:"Data"`
}

type CovidCase struct {
	// date time cannot be parsed
	ConfirmDate    string  `json:"ConfirmDate"`
	No             *int    `json:"No"`
	Age            int     `json:"Age"`
	Gender         *string `json:"Gender"`
	GenderEn       *string `json:"GenderEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       string  `json:"Province"`
	ProvinceId     int     `json:"ProvinceId"`
	ProviceEn      string  `json:"ProvinceEn"`
	District       *string `json:"District"`
	StatQuarantine int     `json:"StatQuarantine"`
}

type CovidSummary struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
