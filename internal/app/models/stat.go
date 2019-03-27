package models

import "time"

//Stat definition
type Stat struct {
	AvgScore    float64       `json:"avgscore"`
	AvgDuration time.Duration `json:"avgduration"`
}
