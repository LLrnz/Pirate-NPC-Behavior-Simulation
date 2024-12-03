package models

import (
	"math/rand"
	"testGUI/pkg/engine"
)

type Simulation struct {
	Env     string
	Objects []engine.Object
}

func (s *Simulation) AddAgent() {
	x := rand.Intn(20) * 16
	y := rand.Intn(20) * 16

	s.Objects = append(s.Objects, NewPirateship(*engine.NewPosition(x, y)))
}
