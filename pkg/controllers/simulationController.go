package controllers

import (
	"fmt"
	"simulationPirate/pkg/models"
	"simulationPirate/pkg/views"
)

type SimulationController struct {
	simu        *models.Simulation
	View        *views.SimulationView
	ControlView *views.ControlView
	events      chan views.Events
}

func NewSimulationController(simu *models.Simulation, view *views.SimulationView, controlView *views.ControlView, events chan views.Events) *SimulationController {
	controller := SimulationController{simu, view, controlView, events}

	//event listener
	go func() {
		for {
			select {
			case evt := <-events:
				if evt == views.AddAgent {
					controller.AddAgent()
				}
			}
		}
	}()

	return &controller
}

func (s *SimulationController) Update(simulation *models.Simulation) {
	s.simu = simulation
	s.View.Update(s.simu)
}

func (s *SimulationController) AddAgent() {
	s.simu.AddAgent()
	fmt.Println("add agent")
}
