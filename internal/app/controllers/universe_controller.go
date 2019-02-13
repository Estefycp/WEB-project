package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

var globalUniverse models.Universe = models.Universe{
	NextID: 0,
	Player: map[int64]*models.Player{},
	Radius: 300.0,
}

// SendUniverse to the client
func SendUniverse(w http.ResponseWriter, r *http.Request) {
	universe := GetUniverse()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(universe)

}

// PostNewPlayer to the client
func PostNewPlayer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	universe := GetUniverse()
	player := CreateAndAddPlayer(universe, r.Form["name"][0])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// PutPlayer and update its values
func PutPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r.ParseForm()

	universe := GetUniverse()
	intid, _ := strconv.Atoi(vars["id"])
	id := int64(intid)
	up, _ := strconv.ParseBool(r.Form["up"][0])
	down, _ := strconv.ParseBool(r.Form["down"][0])
	left, _ := strconv.ParseBool(r.Form["left"][0])
	right, _ := strconv.ParseBool(r.Form["right"][0])

	player := UpdateUniverse(universe, id, up, down, left, right)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// GetUniverse struct
func GetUniverse() *models.Universe {
	return &globalUniverse
}

// CreateAndAddPlayer to universe
func CreateAndAddPlayer(universe *models.Universe, name string) *models.Player {
	player := CreatePlayer(name, universe)
	universe.NextID++
	AddPlayer(universe, &player)
	return &player
}

// UpdateUniverse with a player action
func UpdateUniverse(universe *models.Universe, id int64, up, down, left, right bool) *models.Player {
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
	centerDist := math.Sqrt(math.Pow(player.X+deltaX, 2.0) + math.Pow(player.Y+deltaY, 2.0))
	if centerDist < universe.Radius {
		MovePlayer(player, deltaX, deltaY)
		UpdateRadius(player, deltaR)
	}
	CheckAllCollisions(universe, player)
	return player
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
