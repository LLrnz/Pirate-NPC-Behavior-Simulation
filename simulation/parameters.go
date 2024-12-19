package simulation

import (
	"sync"
	"time"
)

/*
----------------------------------------
************* Parameters ***************
----------------------------------------
*/
// The user will be able to :
// Pause the simulation
// Adjust time speed

type Parameters struct {
	sync.RWMutex
	paused      bool
	step        int
	maxStep     int
	start       time.Time
	maxDuration time.Duration
	timeSpeed   int
}

// NewParameters initializes the Parameters of the sumulation
func NewParameters(maxDuration time.Duration, maxStep int) *Parameters {
	return &Parameters{
		paused:      false,
		step:        0,
		maxStep:     maxStep,
		start:       time.Time{},
		maxDuration: maxDuration,
		timeSpeed:   1,
	}
}

// Getter paused
func (p *Parameters) Paused() bool {
	p.RLock()
	defer p.RUnlock()
	return p.paused
}

// Setter paused
func (p *Parameters) SetPaused(paused bool) {
	p.Lock()
	defer p.Unlock()
	p.paused = paused
}

// Getter step
func (p *Parameters) Step() int {
	p.RLock()
	defer p.RUnlock()
	return p.step
}

// Setter step
func (p *Parameters) AddStep() {
	p.Lock()
	defer p.Unlock()
	p.step = p.step + 1
}

// Getter maxStep
func (p *Parameters) MaxStep() int {
	return p.maxStep
}

// Setter maxStep
func (p *Parameters) SetMaxStep(maxStep int) {
	p.maxStep = maxStep
}

// Getter start
func (p *Parameters) Start() time.Time {
	return p.start
}

// Setter start
func (p *Parameters) SetStart(start time.Time) {
	p.start = start
}

// Getter maxDuration
func (p *Parameters) MaxDuration() time.Duration {
	return p.maxDuration
}

// Setter maxDuration
func (p *Parameters) SetMaxDuration(maxDuration time.Duration) {
	p.maxDuration = maxDuration
}

// Getter timeSpeed
func (p *Parameters) TimeSpeed() int {
	p.RLock()
	defer p.RUnlock()
	return p.timeSpeed
}

// Setter timeSpeed
func (p *Parameters) SetTimeSpeed(timeSpeed int) {
	p.Lock()
	defer p.Unlock()
	p.timeSpeed = timeSpeed
}
