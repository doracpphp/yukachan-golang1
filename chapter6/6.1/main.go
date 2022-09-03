type AutoGenerated []struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDatetime   time.Time `json:"reportDatetime"`
	TimeSeries       []struct {
		TimeDefines []time.Time `json:"timeDefines"`
		Areas       []struct {
			Area struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"area"`
			WeatherCodes []string `json:"weatherCodes"`
			Weathers     []string `json:"weathers"`
			Winds        []string `json:"winds"`
			Waves        []string `json:"waves"`
		} `json:"areas"`
	} `json:"timeSeries"`
	TempAverage struct {
		Areas []struct {
			Area struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"area"`
			Min string `json:"min"`
			Max string `json:"max"`
		} `json:"areas"`
	} `json:"tempAverage,omitempty"`
	PrecipAverage struct {
		Areas []struct {
			Area struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"area"`
			Min string `json:"min"`
			Max string `json:"max"`
		} `json:"areas"`
	} `json:"precipAverage,omitempty"`
}