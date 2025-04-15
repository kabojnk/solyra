package models

// Application represents an API client application
type Application struct {
	ID           string `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
}
