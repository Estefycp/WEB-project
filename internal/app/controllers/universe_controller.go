package controllers

import "../models"

var globalUniverse models.Universe = models.Universe{}

// GetUniverse struct
func GetUniverse() *models.Universe {
	return &globalUniverse
}

// CreateAndAddPlayer to universe
func CreateAndAddPlayer(universe *models.Universe, name string) {
	player := CreatePlayer(name, universe)
	AddPlayer(universe, &player)
}

// UpdateUniverse with a player action
func UpdateUniverse(universe *models.Universe, id int64, up, down, left, right bool) {
	deltaR := 0.001
	deltaX := 0.0
	deltaY := 0.0
	if up {
		deltaY += 1.0
	}
	if down {
		deltaY -= 1.0
	}
	if right {
		deltaX += 1.0
	}
	if left {
		deltaX -= 1.0
	}
	player := universe.Player[id]
	MovePlayer(player, deltaX, deltaY)
	UpdateRadius(player, deltaR)
	CheckAllCollisions(universe, player)
}

// AddPlayer to the universe
func AddPlayer(universe *models.Universe, player *models.Player) {
	universe.Player[player.ID] = player
}

// DeletePlayer from the game
func DeletePlayer(universe *models.Universe, player *models.Player) {
	delete(universe.Player, player.ID)
}

// CheckAllCollisions in the universe
func CheckAllCollisions(universe *models.Universe, player *models.Player) {
	for k1, p1 := range universe.Player {
		if k1 == player.ID {
			continue
		}
		if CheckCollision(p1, player) {
			DeletePlayer(universe, p1)
			DeletePlayer(universe, player)
		}
	}
}
