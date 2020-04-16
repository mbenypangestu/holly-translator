package models

type Location struct {
	ID         string `json:"_id,omitempty"`
	LocationID string `json:"location_id,omitempty"`
	Name       string `json:"name,omitempty"`
}
