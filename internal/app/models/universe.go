package models

// Universe to have all players inside
type Universe struct {
	NextID int64
	Player map[int64]*Player
	Radius float64
}
