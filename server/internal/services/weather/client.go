package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kevinmahoney/etrenank/internal/models"
)

// Client represents a weather API client
type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// WeatherAPIResponse represents the response from WeatherAPI.com
type WeatherAPIResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		LocaltimeEpoch int64   `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		UV         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
		AirQuality struct {
			CO             float64 `json:"co"`
			NO2            float64 `json:"no2"`
			O3             float64 `json:"o3"`
			SO2            float64 `json:"so2"`
			PM25           float64 `json:"pm2_5"`
			PM10           float64 `json:"pm10"`
			USEPAIndex     int     `json:"us-epa-index"`
			GBDEFRAIndex   int     `json:"gb-defra-index"`
		} `json:"air_quality,omitempty"`
	} `json:"current"`
	Astronomy struct {
		Astro struct {
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			MoonPhase        string `json:"moon_phase"`
			MoonIllumination string `json:"moon_illumination"`
		} `json:"astro"`
	} `json:"astronomy"`
}

// NewClient creates a new weather API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://api.weatherapi.com/v1",
	}
}

// GetWeatherByZipCode fetches weather data for a specific zip code
func (c *Client) GetWeatherByZipCode(zipCode string) (*models.WeatherData, *models.AstronomyData, error) {
	url := fmt.Sprintf("%s/forecast.json?key=%s&q=%s&aqi=yes&alerts=no&days=1&astronomy=yes", c.baseURL, c.apiKey, zipCode)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("weather API returned status code %d", resp.StatusCode)
	}
	
	var apiResp WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, nil, err
	}
	
	// Extract weather data
	weatherData := &models.WeatherData{
		CloudCoverPercentage: float64(apiResp.Current.Cloud),
		Humidity:             float64(apiResp.Current.Humidity),
		VisibilityKm:         apiResp.Current.VisKm,
		WindSpeed:            apiResp.Current.WindMph,
		Temperature:          apiResp.Current.TempC,
		Location:             fmt.Sprintf("%s, %s", apiResp.Location.Name, apiResp.Location.Region),
	}
	
	// Add AQI if available
	if apiResp.Current.AirQuality.USEPAIndex > 0 {
		weatherData.AirQualityIndex = float64(apiResp.Current.AirQuality.USEPAIndex)
	}
	
	// Get precipitation for last 24h (this is an approximation from current data)
	weatherData.PrecipitationLast24h = apiResp.Current.PrecipMm
	
	// Extract astronomy data
	// We need to calculate sun altitude based on time of day and location
	// This is a simplified calculation - in a real app, you'd use a proper astronomical library
	
	// For now, we'll use a placeholder value based on whether it's day or night
	var sunAltitude float64
	if apiResp.Current.IsDay == 1 {
		// During day, assume sun is somewhere between 0 and 45 degrees
		sunAltitude = 25.0
	} else {
		// During night, assume sun is below horizon
		sunAltitude = -15.0
	}
	
	// Parse moon illumination as float
	moonIllumination := 0.0
	fmt.Sscanf(apiResp.Astronomy.Astro.MoonIllumination, "%f", &moonIllumination)
	
	astronomyData := &models.AstronomyData{
		SunAltitude:       sunAltitude,
		SunAzimuth:        float64(apiResp.Current.WindDegree), // Using wind direction as a placeholder
		SunriseTime:       apiResp.Astronomy.Astro.Sunrise,
		SunsetTime:        apiResp.Astronomy.Astro.Sunset,
		MoonPhase:         apiResp.Astronomy.Astro.MoonPhase,
		MoonIllumination:  moonIllumination,
	}
	
	return weatherData, astronomyData, nil
}
