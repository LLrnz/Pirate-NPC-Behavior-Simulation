package simulation

import (
	"fmt"
	"sync"
	"time"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"

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
	env        *Environment
	parameters *Parameters
}

func (simu *Simulation) Environment() *Environment {
	return simu.env
}

// Getter Agents returns the list of agents
func (simu *Simulation) Agents() []Agent {
	return simu.Environment().Agents()
}

// Getter Parameters
func (simu *Simulation) Parameters() *Parameters {
	return simu.parameters
}

// NewRandomCoordinates generates random coordinates that is not land
func NewRandomCoordinates() Coordinates {
	for {
		x := rand.Intn(48)
		y := rand.Intn(27)

		// Check if the cell is sea
		if constants.MapGrid[y][x] == 0 {
			return Coordinates{x: x, y: y}
		}

	}
}

// NewSimulation initializes the Simulation
func NewSimulation(maxDuration time.Duration, maxStep int) (simu *Simulation) {
	simu = &Simulation{}
	simu.env = NewEnvironment([]Agent{}, []Harbour{})
	simu.parameters = NewParameters(maxDuration, maxStep)

	// Create harbours with available and sellable resources
	dockedHarbour1 := NewHarbour("Canada", Coordinates{13, 5})
	dockedHarbour1.AddAvailableRessources("Gold")
	dockedHarbour1.AddSellableRessources("Sugar")
	simu.env.AddHarbour(*dockedHarbour1)

	dockedHarbour2 := NewHarbour("Port à Coutollenc", Coordinates{22, 11})
	dockedHarbour2.AddSellableRessources("Gold")
	dockedHarbour2.AddAvailableRessources("Tipiak")
	simu.env.AddHarbour(*dockedHarbour2)

	dockedHarbour3 := NewHarbour("Africa", Coordinates{20, 25})
	dockedHarbour3.AddAvailableRessources("Coffee")
	dockedHarbour3.AddSellableRessources("Tipiak")
	simu.env.AddHarbour(*dockedHarbour3)

	dockedHarbour4 := NewHarbour("Brasil", Coordinates{10, 22})
	dockedHarbour4.AddSellableRessources("Coffee")
	dockedHarbour4.AddAvailableRessources("Sugar")
	simu.env.AddHarbour(*dockedHarbour4)

	// Define nationalities
	nationalities := []constants.Nationality{
		constants.France,
		constants.England,
		constants.Spain,
	}

	// Specify how many merchant ships each nationality should have
	merchantCounts := map[constants.Nationality]int{
		constants.France:  5,
		constants.Spain:   10,
		constants.England: 3,
	}

	privateerCounts := map[constants.Nationality]int{
		constants.France:  2,
		constants.England: 1,
		constants.Spain:   4,
	}

	armyCounts := map[constants.Nationality]int{
		constants.France:  1,
		constants.England: 2,
		constants.Spain:   1,
	}

	numPirates := 1

	// Set up a list of harbours to randomly assign merchant ships to
	merchantHarbours := []*Harbour{dockedHarbour1, dockedHarbour2, dockedHarbour3, dockedHarbour4}

	// Create Pirate Ships (no nationality)
	for i := 0; i < numPirates; i++ {
		id := fmt.Sprintf("Pirate Ship #%d", i)
		coor := NewRandomCoordinates()
		ship := NewShipPirate(id, coor, nil)
		simu.env.AddAgent(ship)
	}

	// Create Merchant Ships per nationality
	for _, nat := range nationalities {
		count := merchantCounts[nat]
		for i := 0; i < count; i++ {
			harbour := merchantHarbours[i%len(merchantHarbours)]
			id := fmt.Sprintf("Merchant Ship [%s] #%d", nat, i)
			ship := NewShipMerchant(id, nat, harbour.Coordinates(), harbour)
			simu.env.AddAgent(ship)
		}
	}

	for _, nat := range nationalities {
		for i := 0; i < privateerCounts[nat]; i++ {
			id := fmt.Sprintf("Privateer Ship [%s] #%d", nat, i)
			coor := NewRandomCoordinates()
			ship := NewShipPrivateer(id, nat, coor, nil)
			simu.env.AddAgent(ship)
		}
	}

	for _, nat := range nationalities {
		for i := 0; i < armyCounts[nat]; i++ {
			id := fmt.Sprintf("Army Ship [%s] #%d", nat, i)
			coor := NewRandomCoordinates()
			ship := NewShipArmy(id, nat, coor, nil)
			simu.env.AddAgent(ship)
		}
	}

	// Special ships (no nationality)
	zoro := NewShipZoro("Zoro", NewRandomCoordinates(), nil)
	simu.env.AddAgent(zoro)

	rackham := NewShipRackham("Rackham", NewRandomCoordinates(), nil)
	simu.env.AddAgent(rackham)

	kraken := NewKraken("Kraken", Coordinates{x: 13, y: 12})
	simu.env.AddAgent(kraken)

	return simu
}

// Run Simulation
func (simu *Simulation) Run() {
	// Start the log microservice
	go simu.Log()

	// Save the simulation's start time using the setter
	simu.Parameters().SetStart(time.Now())

	// Simulation loop
Loop:
	for {
		var wg sync.WaitGroup

		// Increment the step using the getter and setter
		simu.Parameters().AddStep()
		currentStep := simu.Parameters().Step()

		if currentStep == 10 {
			time := simu.Parameters().TimeSpeed()
			simu.Parameters().SetTimeSpeed(time * 10)
		}

		agents := simu.Agents()

		for _, agt := range agents {
			wg.Add(1)

			go func(agt Agent) { // workers
				defer wg.Done()
				agt.Percept(simu.Environment(), currentStep)
				agt.Deliberate(simu.Environment(), currentStep)
				agt.Act(simu.Environment(), currentStep)
			}(agt)
		}
		wg.Wait()

		// Check termination conditions using getters
		if (simu.Parameters().MaxStep() > 0 && simu.Parameters().Step() >= simu.Parameters().MaxStep()) || time.Since(simu.Parameters().Start()) > simu.Parameters().MaxDuration() {
			break Loop
		}

		// Adjust the sleep duration based on timeSpeed using the getter
		time.Sleep(time.Second / (2 * time.Duration(simu.Parameters().TimeSpeed()))) // 2 steps/s default and time speed
	}
}

func (simu *Simulation) Print() {

}

func (simu *Simulation) Log() {

}
