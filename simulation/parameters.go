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
