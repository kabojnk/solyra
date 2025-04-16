package photoquality

import (
	"math"

	"github.com/kevinmahoney/etrenank/internal/models"
)

// CalculateSunriseQuality evaluates the photographic quality of a sunrise/sunset
func CalculateSunriseQuality(weather models.WeatherData, astronomy models.AstronomyData) (float64, map[string]float64, string) {
	// Initialize base score
	qualityScore := 50.0 // Start with neutral score of 50/100

	// Initialize factors map
	factors := make(map[string]float64)

	// === CLOUD COVER ANALYSIS ===
	cloudCover := weather.CloudCoverPercentage
	var cloudScore float64

	// Optimal cloud cover is between 30-70%
	if cloudCover >= 30 && cloudCover <= 70 {
		// Parabolic function peaking at 50% cloud cover
		cloudScore = 25 - 0.02*math.Pow(cloudCover-50, 2)
	} else if cloudCover < 30 {
		// Less dramatic with too few clouds
		cloudScore = cloudCover * 0.6
	} else { // > 70%
		// Too many clouds blocks light
		cloudScore = math.Max(0, 25-(cloudCover-70)*0.8)
	}
	factors["cloud_score"] = cloudScore

	// === ATMOSPHERIC CLARITY ===
	humidity := weather.Humidity
	visibility := weather.VisibilityKm
	aqi := weather.AirQualityIndex
	if aqi == 0 {
		aqi = 50 // Default if not available
	}

	// Humidity factor (40-70% is ideal)
	var humidityScore float64
	if humidity >= 40 && humidity <= 70 {
		humidityScore = 15
	} else if humidity < 40 {
		humidityScore = humidity * 0.3 // Too dry = less dramatic colors
	} else { // > 70%
		humidityScore = math.Max(0, 15-(humidity-70)*0.3) // Too humid = hazy
	}
	factors["humidity_score"] = humidityScore

	// Visibility factor
	visibilityScore := math.Min(15, visibility*1.5)
	factors["visibility_score"] = visibilityScore

	// Air quality factor (moderate pollution can enhance colors)
	var aqiScore float64
	if aqi >= 50 && aqi <= 100 {
		aqiScore = 10 // Peak score for moderate AQI
	} else if aqi < 50 {
		aqiScore = aqi * 0.2 // Gradually increase score up to 50
	} else { // > 100
		aqiScore = math.Max(0, 10-(aqi-100)*0.1) // Decrease score for high pollution
	}
	factors["air_quality_score"] = aqiScore

	// === WEATHER CONDITIONS ===
	recentRain := weather.PrecipitationLast24h
	windSpeed := weather.WindSpeed

	// Recent light rain is good (clears air)
	var rainScore float64
	if recentRain > 0 && recentRain < 5 {
		rainScore = 5
	} else if recentRain >= 5 {
		rainScore = math.Max(0, 5-(recentRain-5)*0.5)
	} else {
		rainScore = 0
	}
	factors["recent_rain_score"] = rainScore

	// Light wind is good (5-15 mph ideal)
	var windScore float64
	if windSpeed >= 5 && windSpeed <= 15 {
		windScore = 5
	} else if windSpeed < 5 {
		windScore = windSpeed * 0.8
	} else { // > 15
		windScore = math.Max(0, 5-(windSpeed-15)*0.3)
	}
	factors["wind_score"] = windScore

	// === CALCULATE FINAL SCORE ===
	qualityScore += cloudScore +
		humidityScore +
		visibilityScore +
		aqiScore +
		rainScore +
		windScore

	// Clamp final score between 0-100
	qualityScore = math.Max(0, math.Min(100, qualityScore))

	return qualityScore, factors, interpretScore(qualityScore)
}

// interpretScore provides a human-readable interpretation of the quality score
func interpretScore(score float64) string {
	if score >= 80 {
		return "Exceptional conditions for dramatic sunrise/sunset photography"
	} else if score >= 65 {
		return "Very good conditions, expect vibrant colors"
	} else if score >= 50 {
		return "Good conditions, some color expected"
	} else if score >= 35 {
		return "Fair conditions, limited color possible"
	} else {
		return "Poor conditions, minimal color expected"
	}
}
