package main

import (
	"sync"

	"github.com/troian/healthcheck"
)

type healthChecks struct {
	healthLock      sync.Mutex
	livenessChecks  map[string]healthcheck.Check
	readinessChecks map[string]healthcheck.Check
}

func NewHealthChecks() *healthChecks {
	health := &healthChecks{
		livenessChecks:  make(map[string]healthcheck.Check),
		readinessChecks: make(map[string]healthcheck.Check),
	}
	return health
}

func (t *healthChecks) AddLivenessCheck(name string, check healthcheck.Check) error {
	t.healthLock.Lock()
	defer t.healthLock.Unlock()

	t.livenessChecks[name] = check

	return nil
}
func (t *healthChecks) AddReadinessCheck(name string, check healthcheck.Check) error {
	t.healthLock.Lock()
	defer t.healthLock.Unlock()

	t.readinessChecks[name] = check

	return nil
}
func (t *healthChecks) RemoveLivenessCheck(name string) error {
	t.healthLock.Lock()
	defer t.healthLock.Unlock()

	delete(t.livenessChecks, name)

	return nil
}
func (t *healthChecks) RemoveReadinessCheck(name string) error {
	t.healthLock.Lock()
	defer t.healthLock.Unlock()

	delete(t.readinessChecks, name)

	return nil
}
