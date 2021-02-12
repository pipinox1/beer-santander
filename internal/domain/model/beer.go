package model

import "time"

type MeetUp struct {
	TotalGuests int64     `json:"total_guests"`
	Location    Location  `json:"location"`
	Date        time.Time `json:"date"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
