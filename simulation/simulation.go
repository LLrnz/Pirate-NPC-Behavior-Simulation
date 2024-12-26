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
		x := rand.Intn(96)
		y := rand.Intn(54)

		// Check if the cell is sea
		if constants.MapGrid[y][x] == 0 {
			return Coordinates{x: x, y: y}
		}

	}
}

// NewSimulation initializes the Simulation
func NewSimulation(maxDuration time.Duration, maxStep int) (simu *Simulation) {
	simu = &Simulation{}
	simu.env = NewEnvironment([]Agent{}, []*Harbour{})
	simu.parameters = NewParameters(maxDuration, maxStep)

	// Haiti (Coordinates{4, 39})
	dockedHarbourHaiti := NewHarbour("Haiti", Coordinates{4, 39})
	// Removed Gold (most valuable) from availability
	dockedHarbourHaiti.AddAvailableRessources("Tobacco")
	dockedHarbourHaiti.AddAvailableRessources("Tipiak")
	// Keep lesser-value resources as sellable
	dockedHarbourHaiti.AddSellableRessources("Spices")
	dockedHarbourHaiti.AddSellableRessources("Sugar")
	dockedHarbourHaiti.AddSellableRessources("Cotton")
	simu.env.AddHarbour(dockedHarbourHaiti)

	// Peru (Coordinates{17, 32})
	dockedHarbourPeru := NewHarbour("Peru", Coordinates{17, 32})
	dockedHarbourPeru.AddAvailableRessources("Silver")
	dockedHarbourPeru.AddAvailableRessources("Cotton")
	dockedHarbourPeru.AddSellableRessources("Cocoa")
	dockedHarbourPeru.AddSellableRessources("Tobacco")
	simu.env.AddHarbour(dockedHarbourPeru)

	// Canada (Coordinates{21, 18})
	dockedHarbourCanada := NewHarbour("Canada", Coordinates{21, 18})
	dockedHarbourCanada.AddAvailableRessources("Coffee")
	dockedHarbourCanada.AddAvailableRessources("Cocoa")
	// Removed Ivory (most valuable) from availability
	dockedHarbourCanada.AddSellableRessources("Sugar")
	dockedHarbourCanada.AddSellableRessources("Rum")
	dockedHarbourCanada.AddSellableRessources("Tipiak")
	simu.env.AddHarbour(dockedHarbourCanada)

	// Argentine (Coordinates{24, 39})
	dockedHarbourArgentine := NewHarbour("Argentine", Coordinates{24, 39})
	dockedHarbourArgentine.AddAvailableRessources("Spices")
	dockedHarbourArgentine.AddAvailableRessources("Weapons")
	// Pearls (most valuable) now only available here
	dockedHarbourArgentine.AddAvailableRessources("Pearls")
	// Removed Gold from sellable
	dockedHarbourArgentine.AddSellableRessources("Tipiak")
	simu.env.AddHarbour(dockedHarbourArgentine)

	// Groenland (Coordinates{37, 8})
	dockedHarbourGroenland := NewHarbour("Groenland", Coordinates{37, 8})
	dockedHarbourGroenland.AddAvailableRessources("Sugar")
	// Removed Silk (most valuable) from availability
	dockedHarbourGroenland.AddSellableRessources("Silver")
	dockedHarbourGroenland.AddSellableRessources("Weapons")
	simu.env.AddHarbour(dockedHarbourGroenland)

	// Compiègne (Coordinates{42, 22})
	dockedHarbourFrance := NewHarbour("Compiègne", Coordinates{42, 22})
	dockedHarbourFrance.AddAvailableRessources("Rum")
	// Removed Pearls (most valuable) from availability
	// Removed Silk and Ivory from sellable
	dockedHarbourFrance.AddSellableRessources("Coffee")
	simu.env.AddHarbour(dockedHarbourFrance)

	// Afrique du Sud (Coordinates{49, 40})
	dockedHarbourAfriqueDuSud := NewHarbour("Afrique du Sud", Coordinates{49, 40})
	dockedHarbourAfriqueDuSud.AddAvailableRessources("Coffee")
	dockedHarbourAfriqueDuSud.AddAvailableRessources("Cotton")
	// Removed Pearls and Ivory from sellable
	dockedHarbourAfriqueDuSud.AddSellableRessources("Tobacco")
	simu.env.AddHarbour(dockedHarbourAfriqueDuSud)

	// Russie (Coordinates{57, 14})
	dockedHarbourRussie := NewHarbour("Russie", Coordinates{57, 14})
	dockedHarbourRussie.AddAvailableRessources("Spices")
	dockedHarbourRussie.AddAvailableRessources("Tipiak")
	dockedHarbourRussie.AddAvailableRessources("Tobacco")
	// Removed Gold (most valuable) from sellable
	dockedHarbourRussie.AddSellableRessources("Cocoa")
	simu.env.AddHarbour(dockedHarbourRussie)

	// Inde (Coordinates{64, 30})
	dockedHarbourInde := NewHarbour("Inde", Coordinates{64, 30})
	dockedHarbourInde.AddAvailableRessources("Sugar")
	// Ivory (most valuable) now only available here
	dockedHarbourInde.AddAvailableRessources("Ivory")
	dockedHarbourInde.AddAvailableRessources("Rum")
	dockedHarbourInde.AddSellableRessources("Silver")
	simu.env.AddHarbour(dockedHarbourInde)

	// Japon (Coordinates{79, 27})
	dockedHarbourJapon := NewHarbour("Japon", Coordinates{79, 27})
	// Gold (most valuable) only available here
	dockedHarbourJapon.AddAvailableRessources("Gold")
	// Silk (most valuable) only available here
	dockedHarbourJapon.AddAvailableRessources("Silk")
	// Removed Pearls from sellable
	dockedHarbourJapon.AddSellableRessources("Coffee")
	dockedHarbourJapon.AddSellableRessources("Rum")
	dockedHarbourJapon.AddSellableRessources("Weapons")
	simu.env.AddHarbour(dockedHarbourJapon)

	// Australie (Coordinates{75, 42})
	dockedHarbourAustralie := NewHarbour("Australie", Coordinates{75, 42})
	dockedHarbourAustralie.AddAvailableRessources("Silver")
	dockedHarbourAustralie.AddAvailableRessources("Cocoa")
	dockedHarbourAustralie.AddAvailableRessources("Weapons")
	// Removed Silk from sellable
	dockedHarbourAustralie.AddSellableRessources("Spices")
	dockedHarbourAustralie.AddSellableRessources("Cotton")
	simu.env.AddHarbour(dockedHarbourAustralie)

	// Define nationalities
	nationalities := []constants.Nationality{
		constants.France,
		constants.England,
		constants.Spain,
	}

	// Specify how many merchant ships each nationality should have
	merchantCounts := map[constants.Nationality]int{
		constants.France:  1,
		constants.Spain:   0,
		constants.England: 0,
	}

	privateerCounts := map[constants.Nationality]int{
		constants.France:  0,
		constants.England: 4,
		constants.Spain:   0,
	}

	armyCounts := map[constants.Nationality]int{
		constants.France:  0,
		constants.England: 0,
		constants.Spain:   0,
	}

	numPirates := 0

	// Set up a list of harbours to randomly assign merchant ships to
	merchantHarbours := []*Harbour{
		dockedHarbourHaiti,
		dockedHarbourPeru,
		dockedHarbourCanada,
		dockedHarbourArgentine,
		dockedHarbourGroenland,
		dockedHarbourFrance,
		dockedHarbourAfriqueDuSud,
		dockedHarbourRussie,
		dockedHarbourInde,
		dockedHarbourJapon,
		dockedHarbourAustralie,
	}

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
			coor := NewRandomCoordinates()
			id := fmt.Sprintf("Merchant Ship [%s] #%d", nat, i)
			ship := NewShipMerchant(id, nat, coor, nil)
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
	//zoro := NewShipZoro("Zoro", NewRandomCoordinates(), nil)
	//simu.env.AddAgent(zoro)

	//rackham := NewShipRackham("Rackham", NewRandomCoordinates(), nil)
	//simu.env.AddAgent(rackham)

	kraken := NewKraken("Kraken", Coordinates{x: 13, y: 12})
	simu.env.AddAgent(kraken)

	print(merchantHarbours)

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

		// Display resource prices and fluctuations at the end of the step
		/*fmt.Printf("\n--- Step %d Resource Prices ---\n", currentStep)
		for resName := range RessourceMap {
			price, fluctuation := simu.Environment().Market().RessourcePrice(simu.Parameters().Step(), RessourceMap[resName])
			fmt.Printf("%s: Price = %v, Fluctuation = %+.2f%%\n", resName, price, fluctuation)
		}*/

		// Adjust the sleep duration based on timeSpeed using the getter
		time.Sleep(time.Second / (1 * time.Duration(simu.Parameters().TimeSpeed()))) // 2 steps/s default and time speed
	}
}

func (simu *Simulation) Print() {

}

func (simu *Simulation) Log() {
	fmt.Print("starting")
}
