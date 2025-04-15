package models

// SunsetQuality represents the quality of a sunset for photography
type SunsetQuality struct {
	ZipCode         string             `json:"zip_code"`
	OverallQuality  float64            `json:"overall_quality"`
	Factors         map[string]float64 `json:"factors"`
	Interpretation  string             `json:"interpretation"`
	WeatherData     WeatherData        `json:"weather_data"`
	AstronomyData   AstronomyData      `json:"astronomy_data"`
	LastUpdated     string             `json:"last_updated"`
	ExpiresAt       string             `json:"expires_at"`
}

// WeatherData contains meteorological information from weather APIs
type WeatherData struct {
	CloudCoverPercentage  float64 `json:"cloud_cover_percentage"`
	Humidity              float64 `json:"humidity"`
	VisibilityKm          float64 `json:"visibility_km"`
	AirQualityIndex       float64 `json:"air_quality_index"`
	PrecipitationLast24h  float64 `json:"precipitation_last_24h"`
	WindSpeed             float64 `json:"wind_speed"`
	Temperature           float64 `json:"temperature"`
	Location              string  `json:"location"`
}

// AstronomyData contains sun/moon position information
type AstronomyData struct {
	SunAltitude    float64 `json:"sun_altitude"`
	SunAzimuth     float64 `json:"sun_azimuth"`
	SunriseTime    string  `json:"sunrise_time"`
	SunsetTime     string  `json:"sunset_time"`
	MoonPhase      string  `json:"moon_phase"`
	MoonIllumination float64 `json:"moon_illumination"`
}
