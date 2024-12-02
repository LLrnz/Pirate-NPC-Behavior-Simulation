package simulation

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

/*
---------------------------------------
************* Simulation **************
---------------------------------------
*/
// The simulation handles the environment and has parameters
// It is responsible for the simulation loop where our agents can interact

type Simulation struct {
	env        Environment
	agents     []Agent
	objects    []Object
	parameters Parameters
}

func (simu *Simulation) Environment() *Environment {
	return &simu.env
}

// Agents returns the list of agents
func (simu *Simulation) Agents() []Agent {
	simu.Environment().agentLock.RLock()
	defer simu.Environment().agentLock.RUnlock()
	return simu.Environment().Agents()
}

// NewRandomCoordinates generates random coordinates
func NewRandomCoordinates() Coordinates {
	return Coordinates{
		x: rand.Intn(800),
		y: rand.Intn(600),
	}
}

// NewSimulation initializes the Simulation [[professor code]]
func NewSimulation(maxDuration time.Duration, maxStep int) (simu *Simulation) {
	simu = &Simulation{}
	simu.agents = make([]Agent, 0)
	simu.objects = make([]Object, 0)

	simu.env = *NewEnvironment([]Agent{}, []Object{})

	simu.parameters = Parameters{
		maxDuration: maxDuration,
		maxStep:     maxStep,
		timeSpeed:   1,
	}

	docked := false
	// création des agents
	for i := 0; i < 200; i++ {
		// création de l'agent
		id := fmt.Sprintf("Ship #%d", i)
		coor := NewRandomCoordinates()
		ship := NewShipMerchant(id, 10, coor, docked, &simu.env)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
		docked = false
	}

	// Kraken
	kraken := NewKraken("Kraken #1", 200, 10, 4, Coordinates{x: 100, y: 100})
	simu.agents = append(simu.agents, kraken)
	simu.env.AddAgent(kraken)

	return simu
}

// Run Simulation [[professor code]]
func (simu *Simulation) Run() {
	// Démarrage du micro-service de Log
	go simu.Log()
	// Démarrage du micro-service d'affichage
	//go simu.Print()

	// On sauvegarde la date du début de la simulation
	simu.parameters.start = time.Now()

	// Boucle de simulation
Loop:
	for {
		var wg sync.WaitGroup
		simu.parameters.step += 1

		simu.Environment().agentLock.Lock()
		for _, agt := range simu.agents {
			wg.Add(1)

			go func(agt Agent) { // workers
				defer wg.Done()
				agt.Percept(&simu.env)
				agt.Deliberate()
				agt.Act(&simu.env)

			}(agt)

		}
		wg.Wait()

		if (simu.parameters.maxStep > 0 && simu.parameters.step >= simu.parameters.maxStep) || time.Since(simu.parameters.start) > simu.parameters.maxDuration {
			fmt.Println()
			break Loop
		}

		simu.Environment().agentLock.Unlock()

		time.Sleep(time.Second / 50)

	}

}

func (simu *Simulation) Print() {

}

func (simu *Simulation) Log() {

}
