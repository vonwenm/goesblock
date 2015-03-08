// Package allows store errors count from external services.
// Useful when any of external services is down / busy / answers with timeout and you need to slow down work for some period.

package goesblock

import (
	"time"
)

var (
	cleanInterval = time.Second * 30 // rewrite counters every 30 seconds
)

func init() {
	ticker := time.Tick(cleanInterval)

	go func() {
		for _ = range ticker {
			services.clean()
		}
	}()
}
