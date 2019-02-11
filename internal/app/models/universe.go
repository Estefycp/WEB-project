package models

// Universe to have all players inside
type Universe struct {
	NextID int64             `json:"next"`
	Player map[int64]*Player `json:"player"`
	Radius float64           `json:"radius"`
}
