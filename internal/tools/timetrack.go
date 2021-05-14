package tools

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time) string {
	elapsed := time.Since(start)
	return fmt.Sprintf("time : %s", elapsed)
}
