package controllers

import (
	"fmt"
	"time"

	"gitlab.utc.fr/arsaintm/ia04-projet/UI/widgets"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

type SimulationController struct {
	simu        *simulation.Simulation
	View        *widgets.SimulationView
	ControlView *widgets.ControlView
	events      chan widgets.Events
	refreshRate float64
}

func NewSimulationController(simu *simulation.Simulation, view *widgets.SimulationView, controlView *widgets.ControlView, events chan widgets.Events, refreshRate float64) *SimulationController {
	controller := SimulationController{simu, view, controlView, events, refreshRate}

	//event listener
	go func() {
		for {
			select {
			case evt := <-events:
				if evt == widgets.AddAgent {
					controller.AddAgent()
				}
			}
		}
	}()

	return &controller
}

func (s *SimulationController) Start() {
	go func() {
		dur := time.Duration(float64(time.Second) * (1 / s.refreshRate))
		time.Sleep(dur)
		s.View.Update(s.simu)
	}()
}

func (s *SimulationController) AddAgent() {
	//s.simu.AddAgent()
	fmt.Println("add agent")
}
