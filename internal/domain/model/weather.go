package model

type Weather struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				MaxtempC float64 `json:"maxtemp_c"`
				MintempC float64 `json:"mintemp_c"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
