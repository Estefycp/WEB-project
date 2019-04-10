package routines

import (
	"time"

	"github.com/Estefycp/controllers"
)

// schedule a task with a constant interval.
func schedule(what func(), delay time.Duration) chan bool {
	ticker := time.NewTicker(delay)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				what()
			}
		}
	}()
	return done
}

// StartRoutines for the server
func StartRoutines() {
	schedule(controllers.DeleteInactiveRoutine, time.Minute)
}
