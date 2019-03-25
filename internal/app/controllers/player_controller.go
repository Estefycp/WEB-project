package controllers

import (
	"math"
	"math/rand"
	"time"

	"../models"
)

// CreatePlayer for the game
func CreatePlayer(name string, universe *models.Universe) models.Player {
	initialRadius := 1.0

	randR := rand.Float64() * universe.Radius
	randA := rand.Float64() * 2 * math.Pi

	initalX := randR * math.Cos(randA)
	initalY := randR * math.Sin(randA)

	var initialScore int64
	initialScore = 100

	return models.Player{
		ID:       universe.NextID,
		Name:     name,
		Score:    initialScore,
		Radius:   initialRadius,
		X:        initalX,
		Y:        initalY,
		LastMove: time.Now(),
	}
}

// MovePlayer by the given parameters
func MovePlayer(player *models.Player, deltaX, deltaY float64) {
	player.X += deltaX
	player.Y += deltaY
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
