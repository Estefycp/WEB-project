package controllers

import (
	"math"
	"math/rand"
	"time"

	"../models"
)

// CreatePlayer for the game
func CreatePlayer(name string, skin int64, universe *models.Universe) models.Player {
	initialRadius := 1.0

	randR := rand.Float64() * universe.Radius
	randA := rand.Float64() * 2 * math.Pi

	initalX := randR * math.Cos(randA)
	initalY := randR * math.Sin(randA)

	var initialScore int64
	initialScore = 100

	return models.Player{
		ID:          universe.NextID,
		Name:        name,
		Score:       initialScore,
		Radius:      initialRadius,
		X:           initalX,
		Y:           initalY,
		LastMove:    time.Now(),
		Born:        time.Now(),
		Skin:        skin,
		Vx:          0,
		Vy:          0,
		RightCharge: 0,
		LeftCharge:  0,
		UpCharge:    0,
		DownCharge:  0,
	}
}

// MovePlayer by the given parameters
func MovePlayer(player *models.Player, deltaX, deltaY float64) {
	player.X += deltaX
	player.Y += deltaY
}

// LaunchPlayer by their velocity
func LaunchPlayer(player *models.Player) {
	player.X += player.Vx
	if player.Vx > 1.0 {
		player.Vx -= 1.0
	} else if player.Vx < -1.0 {
		player.Vx += 1.0
	} else {
		player.Vx = 0.0
	}

	player.Y += player.Vy
	if player.Vy > 1.0 {
		player.Vy -= 1.0
	} else if player.Vy < -1.0 {
		player.Vy += 1.0
	} else {
		player.Vy = 0.0
	}
}

// UpdateRadius of the player by the given amount
func UpdateRadius(player *models.Player, deltaR float64) {
	player.Radius += deltaR
}

// UpdateScore of the player to the current one
func UpdateScore(player *models.Player) {
	player.Score = int64(player.Radius * 100.0)
}

// CheckCollision between players
func CheckCollision(p1, p2 *models.Player) bool {
	distance := math.Sqrt(math.Pow(p1.X-p2.X, 2.0) + math.Pow(p1.Y-p2.Y, 2.0))
	minDistante := p1.Radius + p2.Radius
	return distance < minDistante
}
