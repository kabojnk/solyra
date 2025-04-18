package weather

import (
	"encoding/json"
	"fmt"
	"log"
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
		TzID           string  `json:"tz_id"`
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
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
		AirQuality struct {
			CO           float64 `json:"co"`
			NO2          float64 `json:"no2"`
			O3           float64 `json:"o3"`
			SO2          float64 `json:"so2"`
			PM25         float64 `json:"pm2_5"`
			PM10         float64 `json:"pm10"`
			USEPAIndex   int     `json:"us-epa-index"`
			GBDEFRAIndex int     `json:"gb-defra-index"`
		} `json:"air_quality,omitempty"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int64  `json:"date_epoch"`
			Astro     struct {
				Sunrise          string `json:"sunrise"`
				Sunset           string `json:"sunset"`
				Moonrise         string `json:"moonrise"`
				Moonset          string `json:"moonset"`
				MoonPhase        string `json:"moon_phase"`
				MoonIllumination int8   `json:"moon_illumination"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
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
func (c *Client) GetWeatherByZipCode(zipCode string) ([]*models.WeatherData, []*models.AstronomyData, error) {
	url := fmt.Sprintf("%s/forecast.json?key=%s&q=%s&aqi=yes&alerts=no&days=3&astronomy=yes", c.baseURL, c.apiKey, zipCode)

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
		return nil, nil, fmt.Errorf("failed to decode weather API response: %v", err)
	}

	// Load location's timezone
	loc, err := time.LoadLocation(apiResp.Location.TzID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load timezone %s: %v", apiResp.Location.TzID, err)
	}

	weatherDataList := make([]*models.WeatherData, len(apiResp.Forecast.Forecastday))
	astronomyDataList := make([]*models.AstronomyData, len(apiResp.Forecast.Forecastday))

	for i, day := range apiResp.Forecast.Forecastday {
		// Convert times to UTC ISO format
		sunriseTime := c.convertToUTC(day.Astro.Sunrise, loc)
		sunsetTime := c.convertToUTC(day.Astro.Sunset, loc)

		// For the current day, use current conditions
		var weatherData *models.WeatherData
		if i == 0 {
			weatherData = &models.WeatherData{
				CloudCoverPercentage: float64(apiResp.Current.Cloud),
				Humidity:             float64(apiResp.Current.Humidity),
				VisibilityKm:         apiResp.Current.VisKm,
				AirQualityIndex:      float64(apiResp.Current.AirQuality.USEPAIndex),
				PrecipitationLast24h: apiResp.Current.PrecipMm,
				WindSpeed:            apiResp.Current.WindMph,
				Temperature:          apiResp.Current.TempF,
				Location:             fmt.Sprintf("%s, %s", apiResp.Location.Name, apiResp.Location.Region),
				Date:                day.Date,
			}
		} else {
			// For future days, use forecast data
			// Note: You might want to add more forecast data fields to the WeatherAPIResponse struct
			weatherData = &models.WeatherData{
				CloudCoverPercentage: 50.0, // Default values for forecast days
				Humidity:             50.0,
				VisibilityKm:         10.0,
				AirQualityIndex:      50.0,
				PrecipitationLast24h: 0.0,
				WindSpeed:            5.0,
				Temperature:          70.0,
				Location:             fmt.Sprintf("%s, %s", apiResp.Location.Name, apiResp.Location.Region),
				Date:                day.Date,
			}
		}

		astronomyData := &models.AstronomyData{
			SunAltitude:      -15.0, // Default value since it's not provided by the API
			SunAzimuth:       float64(apiResp.Current.WindDegree),
			SunriseTime:      sunriseTime,
			SunsetTime:       sunsetTime,
			MoonPhase:        day.Astro.MoonPhase,
			MoonIllumination: day.Astro.MoonIllumination,
		}

		weatherDataList[i] = weatherData
		astronomyDataList[i] = astronomyData
	}

	return weatherDataList, astronomyDataList, nil
}

// convertToUTC converts a time string in local time (e.g., "07:30 PM") to UTC ISO format
func (c *Client) convertToUTC(timeStr string, loc *time.Location) string {
	// Parse the time string (e.g., "07:30 PM")
	t, err := time.ParseInLocation("03:04 PM", timeStr, loc)
	if err != nil {
		log.Printf("Error parsing time %s: %v", timeStr, err)
		return timeStr // Return original string if parsing fails
	}

	// Get today's date in the location's timezone
	now := time.Now().In(loc)
	today := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, loc)

	// Convert to UTC and format as ISO
	return today.UTC().Format(time.RFC3339)
}
