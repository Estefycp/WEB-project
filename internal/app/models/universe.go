package models

// Universe to have all players inside
type Universe struct {
	player map[int64]Player
	radius float64
}
