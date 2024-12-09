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
	simu.agents = make([]Agent, 0)
	simu.objects = make([]Object, 0)

	simu.env = *NewEnvironment([]Agent{}, []Object{}, []Harbour{})

	simu.parameters = Parameters{
		maxDuration: maxDuration,
		maxStep:     maxStep,
		timeSpeed:   1,
	}

	// Create the harbour
	dockedHarbour1 := NewHarbour("Canada", Coordinates{13, 5})
	simu.env.AddHarbour(*dockedHarbour1)
	dockedHarbour1.AddAvailableRessources("Gold")
	dockedHarbour1.AddSellableRessources("Sugar")

	dockedHarbour2 := NewHarbour("Port à Coutollenc", Coordinates{22, 11})
	simu.env.AddHarbour(*dockedHarbour2)
	dockedHarbour2.AddSellableRessources("Gold")
	dockedHarbour2.AddAvailableRessources("Tipiak")

	// Create the harbour
	dockedHarbour3 := NewHarbour("Africa", Coordinates{20, 25})
	simu.env.AddHarbour(*dockedHarbour3)
	dockedHarbour3.AddAvailableRessources("Coffee")
	dockedHarbour3.AddSellableRessources("Tipiak")

	dockedHarbour4 := NewHarbour("Brasil", Coordinates{10, 22})
	simu.env.AddHarbour(*dockedHarbour4)
	dockedHarbour4.AddSellableRessources("Coffee")
	dockedHarbour4.AddAvailableRessources("Sugar")

	// Nationalities to choose from for ships
	nationalities := []constants.Nationality{
		constants.France,
		constants.England,
		constants.Spain,
	}

	// Number of ships to create for each type
	numPirate := 0
	numMerchant := 1
	numPrivateer := 0
	numArmy := 0

	// Create Pirate Ships
	for i := 0; i < numPirate; i++ {
		id := fmt.Sprintf("Pirate Ship #%d", i)
		coor := NewRandomCoordinates()
		ship := NewShipPirate(id, coor, nil)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create Merchant Ships
	for i := 0; i < numMerchant; i++ {
		id := fmt.Sprintf("Merchant 1 Ship #%d", i)
		nat := nationalities[i%len(nationalities)]
		ship := NewShipMerchant(id, nat, dockedHarbour1.Coordinates(), dockedHarbour1)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create Merchant Ships
	for i := 0; i < numMerchant; i++ {
		id := fmt.Sprintf("Merchant 2 Ship #%d", i)
		nat := nationalities[i%len(nationalities)]
		ship := NewShipMerchant(id, nat, dockedHarbour4.Coordinates(), dockedHarbour4)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create Merchant Ships
	for i := 0; i < numMerchant; i++ {
		id := fmt.Sprintf("Merchant 3 Ship #%d", i)
		nat := nationalities[i%len(nationalities)]
		ship := NewShipMerchant(id, nat, dockedHarbour3.Coordinates(), dockedHarbour3)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create Merchant Ships
	for i := 0; i < numMerchant; i++ {
		id := fmt.Sprintf("Merchant 4 Ship #%d", i)
		nat := nationalities[i%len(nationalities)]
		ship := NewShipMerchant(id, nat, dockedHarbour2.Coordinates(), dockedHarbour2)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create Privateer Ships
	for i := 0; i < numPrivateer; i++ {
		id := fmt.Sprintf("Privateer Ship #%d", i)
		coor := NewRandomCoordinates()
		nat := nationalities[i%len(nationalities)]
		ship := NewShipPrivateer(id, nat, coor, nil)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create Army Ships
	for i := 0; i < numArmy; i++ {
		id := fmt.Sprintf("Army Ship #%d", i)
		coor := NewRandomCoordinates()
		nat := nationalities[i%len(nationalities)]
		ship := NewShipArmy(id, nat, coor, nil)
		simu.agents = append(simu.agents, ship)
		simu.env.AddAgent(ship)
	}

	// Create one Zoro Ship (no nationality)
	zoro := NewShipZoro("Zoro", NewRandomCoordinates(), nil)
	simu.agents = append(simu.agents, zoro)
	simu.env.AddAgent(zoro)

	// Create one Rackham Ship (no nationality)
	rackham := NewShipRackham("Rackham", NewRandomCoordinates(), nil)
	simu.agents = append(simu.agents, rackham)
	simu.env.AddAgent(rackham)

	// Create one Kraken
	kraken := NewKraken("Kraken", Coordinates{x: 13, y: 12})
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
				agt.Percept(&simu.env, simu.parameters.step)
				agt.Deliberate(&simu.env, simu.parameters.step)
				agt.Act(&simu.env, simu.parameters.step)

			}(agt)

		}
		wg.Wait()

		if (simu.parameters.maxStep > 0 && simu.parameters.step >= simu.parameters.maxStep) || time.Since(simu.parameters.start) > simu.parameters.maxDuration {
			fmt.Println()
			break Loop
		}

		simu.Environment().agentLock.Unlock()

		time.Sleep(time.Second / 4)

	}

}

func (simu *Simulation) Print() {

}

func (simu *Simulation) Log() {

}
