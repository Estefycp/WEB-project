package models

import "time"

//Player definition
type Player struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Score    int64     `json:"score"`
	Radius   float64   `json:"radius"`
	X        float64   `json:"x"`
	Y        float64   `json:"y"`
	LastMove time.Time `json:"lastmove"`
	Born     time.Time `json:"born"`
	Skin     int64     `json:"skin"`
}
