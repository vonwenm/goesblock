// Package allows store errors count from external services.
// Useful when any of external services is down / busy / answers with timeout and you need to slow down work for some period.

package goesblock

import (
	"sync/atomic"
)

var (
	blockAfter int64 = 10          // assume that service is down if got 10+ errors in 30 seconds
	services         = &Services{} // variable to store data
)

// Get all described (known) services.
func Get() *Services {
	return services
}

type Services struct {
	ExternalServiceOne service
	ExternalServiceTwo service
}

// Does any of external services is down?
func (e *Services) Down() bool {
	return e.ExternalServiceOne.Down() || e.ExternalServiceTwo.Down()
}

// Clean all counters. Used while time.Tick()
func (e *Services) clean() {
	e.ExternalServiceOne.clean()
	e.ExternalServiceTwo.clean()
}

// Structure stores error counter for single external service.
type service struct {
	counter int64
}

// Increment current service error counter.
func (e *service) IncError() {
	atomic.AddInt64(&e.counter, 1)
}

// Does current service is down?
func (e *service) Down() bool {
	return atomic.LoadInt64(&e.counter) > blockAfter
}

// Clean current service error counter.
func (e *service) clean() {
	atomic.AddInt64(&e.counter, -atomic.LoadInt64(&e.counter))
}
