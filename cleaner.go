package goesblock

import (
	"time"
)

var (
	cleanInterval       = time.Second * 30 // rewrite counters every 30 seconds
	blockAfter    int64 = 10               // assume that service is down if got 10+ errors in 30 seconds
)

func init() {
	ticker := time.Tick(cleanInterval)

	go func() {
		for _ = range ticker {
			services.clean()
		}
	}()
}
