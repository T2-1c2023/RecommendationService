package model

type ProximityRule struct {
	Radius  uint `json:"radius" binding:"required"`
	Enabled bool `json:"enabled" binding:"required"`
}
