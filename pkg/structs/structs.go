package structs

type Location struct {
	Name string `json:"name"`
	Region string `json:"region"`
	Country         string  `json:"country"`
	Latitude        float64 `json:"lat"`
	Longitude       float64 `json:"lon"`
	TimezoneID      string  `json:"tz_id"`
	LocaltimeEpoch  int64   `json:"localtime_epoch"`
	Localtime       string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code string `json:"code"`
}