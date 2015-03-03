package goesblock

import (
	"sync/atomic"
)

var (
	services = &Services{}
)

type Services struct {
	ExternalServiceOne helper
	ExternalServiceTwo helper
}

// Get all described (known) services.
func Service() *Services {
	return services
}

// Does any of external services is down?
func (e *Services) Down() bool {
	return e.ExternalServiceOne.Down() && e.ExternalServiceTwo.Down()
}

// Clean all counters. Used while time.Tick()
func (e *Services) clean() {
	e.ExternalServiceOne.clean()
	e.ExternalServiceTwo.clean()
}

type helper struct {
	counter int64
}

// Increment current service error counter.
func (e *helper) IncError() {
	atomic.AddInt64(&e.counter, 1)
}

// Does current service is down?
func (e *helper) Down() bool {
	return atomic.LoadInt64(&e.counter) > blockAfter
}

// Clean current service error counter.
func (e *helper) clean() {
	atomic.AddInt64(&e.counter, -atomic.LoadInt64(&e.counter))
}
