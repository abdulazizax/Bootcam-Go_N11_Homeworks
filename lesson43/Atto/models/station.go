package models

import "time"

type StationResponse struct {
	Station_ID   string    `json:"station_id"`
	Station_Name string    `json:"station_name"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type StationRequest struct {
	Station_Name string `json:"station_name"`
}
