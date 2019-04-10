package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Estefycp/WEB-project/internal/app/models"
	"github.com/Estefycp/WEB-project/internal/app/storage"
)

// GetStats from the game
func GetStats(w http.ResponseWriter, r *http.Request) {
	scores, errscore := storage.GetInstance().GetScores()
	if errscore != nil {
		json.NewEncoder(w).Encode(errscore)
		return
	}
	durs, errdur := storage.GetInstance().GetAliveDurations()
	if errdur != nil {
		json.NewEncoder(w).Encode(errdur)
		return
	}

	stat := models.Stat{
		AvgScore:    average(scores),
		AvgDuration: averageDur(durs),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stat)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func averageDur(ds []time.Duration) time.Duration {
	// t := time.Now()
	total := 0.0
	for _, d := range ds {
		total += d.Seconds()
	}
	avgd, _ := time.ParseDuration(fmt.Sprintf("%f", (total/float64(len(ds)))) + "s")
	// totald := total.Sub(t)
	// avgd, _ := time.ParseDuration(strconv.Itoa(int(float64(totald.Nanoseconds()) / float64(len(ds)))))
	return avgd
}
