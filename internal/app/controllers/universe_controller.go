package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Estefycp/WEB-project/internal/app/models"
	"github.com/Estefycp/WEB-project/internal/app/storage"
	"github.com/gorilla/mux"
)

var globalUniverse = models.Universe{
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
	player := CreateAndAddPlayer(universe, r.Form["name"][0], r.Form["skin"][0])
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

	// player := UpdateUniverse(universe, id, up, down, left, right)
	player := UpdateChargeUniverse(universe, id, up, down, left, right)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// GetUniverse struct
func GetUniverse() *models.Universe {
	return &globalUniverse
}

// CreateAndAddPlayer to universe
func CreateAndAddPlayer(universe *models.Universe, name string, skin string) *models.Player {
	skinn, err := strconv.ParseInt(skin, 10, 64)
	if err != nil {
		panic(err)
	}
	player := CreatePlayer(name, skinn, universe)
	universe.NextID++
	AddPlayer(universe, &player)
	return &player
}

// UpdateUniverse with a player action
func UpdateUniverse(universe *models.Universe, id int64, up, down, left, right bool) *models.Player {
	player := universe.Player[id]
	deltaR := 0.001
	deltaX := 0.0
	deltaY := 0.0
	if up {
		deltaY += 0.5
	}
	if down {
		deltaY -= 0.5
	}
	if right {
		deltaX += 0.5
	}
	if left {
		deltaX -= 0.5
	}
	centerDist := math.Sqrt(math.Pow(player.X+deltaX, 2.0) + math.Pow(player.Y+deltaY, 2.0))
	if centerDist < universe.Radius {
		MovePlayer(player, deltaX, deltaY)
		UpdateRadius(player, deltaR)
		player.LastMove = time.Now()
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
	storage.GetInstance().SaveScore(player)
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

// DeleteInactive players
func DeleteInactive(universe *models.Universe) {
	minutesToInactive := 1.0
	t := time.Now()
	for _, p := range universe.Player {
		if t.Sub(p.LastMove).Minutes() > minutesToInactive {
			DeletePlayer(universe, p)
		}
	}
}

// DeleteInactiveRoutine on the singleton universe
func DeleteInactiveRoutine() {
	DeleteInactive(GetUniverse())
}

// UpdateChargeUniverse for a player
func UpdateChargeUniverse(universe *models.Universe, id int64, up, down, left, right bool) *models.Player {
	player := universe.Player[id]
	deltaR := 0.001
	if up {
		player.UpCharge++
	} else if player.UpCharge > 0 {
		player.Vy = (float64(player.UpCharge) / 5.0)
		player.UpCharge = 0
	}
	if down {
		player.DownCharge++
	} else if player.DownCharge > 0 {
		player.Vy = -(float64(player.DownCharge) / 5.0)
		player.DownCharge = 0
	}
	if right {
		player.RightCharge++
	} else if player.RightCharge > 0 {
		player.Vx = (float64(player.RightCharge) / 5.0)
		player.RightCharge = 0
	}
	if left {
		player.LeftCharge++
	} else if player.LeftCharge > 0 {
		player.Vx = -(float64(player.LeftCharge) / 5.0)
		player.LeftCharge = 0
	}
	UpdateRadius(player, deltaR)
	player.LastMove = time.Now()
	return player
}

// MoveAll players
func MoveAll(universe *models.Universe) {
	for _, player := range universe.Player {
		centerDist := math.Sqrt(math.Pow(player.X+player.Vx, 2.0) + math.Pow(player.Y+player.Vy, 2.0))
		if centerDist < universe.Radius {
			LaunchPlayer(player)
			CheckAllCollisions(universe, player)
		}
	}
}

// MoveAllRoutine on the singleton universe
func MoveAllRoutine() {
	MoveAll(GetUniverse())
}
