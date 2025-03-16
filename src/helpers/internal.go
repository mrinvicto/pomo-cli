package helpers

import (
	"fmt"
	"time"
)

func Timer(durationInSeconds int) {
	duration := time.Duration(durationInSeconds) * time.Second // Pomodoro session duration
	start := time.Now()

	for {
		elapsed := time.Since(start)
		remaining := duration - elapsed

		if remaining <= 0 {
			fmt.Print("\rTime's up!                          \n")
			break
		}

		// Format time as MM:SS
		mins := int(remaining.Minutes())
		secs := int(remaining.Seconds()) % 60
		fmt.Printf("\rTime Left: %02d:%02d", mins, secs)

		time.Sleep(1 * time.Second)
	}
}
