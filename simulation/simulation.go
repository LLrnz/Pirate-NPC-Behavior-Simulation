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

func (simu *Simulation) CountAgents() map[string]int {
	// Take all agents in the simulation
	agents := simu.Environment().Agents()
	agentCounts := make(map[string]int)

	for _, agent := range agents {
		key := ""
		switch agent.AgtType() {
		case "Ship": // If Ship type transform to ship
			ship, _ := agent.(*Ship)
			if ship.Nationality() != constants.Without {
				// Include nationality in the key for ships with a nationality
				key = fmt.Sprintf("%s[%s]", ship.ShipType(), ship.Nationality())
			} else {
				// Ships without nationality
				key = ship.ShipType()
			}
		default:
			key = agent.AgtType()
		}

		// Add one to the number of this type of agents
		agentCounts[key]++
	}

	return agentCounts
}

// Getter Parameters
func (simu *Simulation) Parameters() *Parameters {
	return simu.parameters
}

// NewRandomCoordinates generates random coordinates that is not land
func NewRandomCoordinates() Coordinates {
	for {
		rand.Seed(uint64(time.Now().UnixNano()))
		x := rand.Intn(96)
		y := rand.Intn(54)

		// Check if the cell is sea
		if constants.MapGrid[y][x] == 0 {
			return Coordinates{x: x, y: y}
		}

	}
}

// Add a kraken to the simulation
func (simu *Simulation) AddKraken() {
	krakenNum := simu.CountAgents()["Kraken"]
	id := fmt.Sprintf("Kraken #%d", krakenNum+1)
	coor := NewRandomCoordinates()
	kraken := NewKraken(id, coor)
	simu.env.AddAgent(kraken)
}

// Add a pirate to the simulation
func (simu *Simulation) AddPirate() {
	pirateNum := simu.CountAgents()["Pirate"]
	id := fmt.Sprintf("Pirate Ship #%d", pirateNum+1)
	coor := NewRandomCoordinates()
	ship := NewShipPirate(id, coor, nil)
	simu.env.AddAgent(ship)
}

// Add an Army ship with nationality to the simulation
func (simu *Simulation) AddArmy(nationality constants.Nationality) {
	num := fmt.Sprintf("Army[%s]", nationality)
	armyNum := simu.CountAgents()[num]
	id := fmt.Sprintf("Army Ship [%s] #%d", nationality, armyNum+1)
	coor := NewRandomCoordinates()
	ship := NewShipArmy(id, nationality, coor, nil)
	simu.env.AddAgent(ship)
}

// Add a Merchant ship with nationality to the simulation
func (simu *Simulation) AddMerchant(nationality constants.Nationality) {
	num := fmt.Sprintf("Merchant[%s]", nationality)
	merchantNum := simu.CountAgents()[num]
	id := fmt.Sprintf("Merchant Ship [%s] #%d", nationality, merchantNum+1)
	coor := NewRandomCoordinates()
	ship := NewShipMerchant(id, nationality, coor, nil)
	simu.env.AddAgent(ship)
}

// Add a Privateer ship with nationality to the simulation
func (simu *Simulation) AddPrivateer(nationality constants.Nationality) {
	num := fmt.Sprintf("Privateer[%s]", nationality)
	privateerNum := simu.CountAgents()[num]
	id := fmt.Sprintf("Privateer Ship [%s] #%d", nationality, privateerNum+1)
	coor := NewRandomCoordinates()
	ship := NewShipPrivateer(id, nationality, coor, nil)
	simu.env.AddAgent(ship)
}

// NewSimulation initializes the Simulation
func NewSimulation(maxDuration time.Duration, maxStep int) (simu *Simulation) {
	simu = &Simulation{}
	simu.env = NewEnvironment([]Agent{}, []*Harbour{})
	simu.parameters = NewParameters(maxDuration, maxStep)

	//-------------------------------------------------------------
	// Californie (Coordinates{12, 24})
	dockedHarbourCalifornie := NewHarbour("Californie", Coordinates{12, 24})
	dockedHarbourCalifornie.AddAvailableRessources("Tobacco")
	dockedHarbourCalifornie.AddAvailableRessources("Cotton")
	dockedHarbourCalifornie.AddSellableRessources("Coffee")
	dockedHarbourCalifornie.AddSellableRessources("Spices")
	dockedHarbourCalifornie.AddSellableRessources("Ivory")
	dockedHarbourCalifornie.AddAvailableRessources("Tipiak")
	simu.env.AddHarbour(dockedHarbourCalifornie)

	//-------------------------------------------------------------
	// Peru (Coordinates{17, 32})
	dockedHarbourPeru := NewHarbour("Peru", Coordinates{17, 32})
	dockedHarbourPeru.AddAvailableRessources("Silver")
	dockedHarbourPeru.AddAvailableRessources("Cotton")
	dockedHarbourPeru.AddSellableRessources("Cocoa")
	dockedHarbourPeru.AddSellableRessources("Tobacco")
	dockedHarbourPeru.AddSellableRessources("Tipiak")
	simu.env.AddHarbour(dockedHarbourPeru)

	//-------------------------------------------------------------
	// Canada (Coordinates{21, 18})
	dockedHarbourCanada := NewHarbour("Canada", Coordinates{21, 18})
	dockedHarbourCanada.AddAvailableRessources("Coffee")
	dockedHarbourCanada.AddAvailableRessources("Cocoa")
	dockedHarbourCanada.AddSellableRessources("Sugar")
	dockedHarbourCanada.AddSellableRessources("Rum")
	dockedHarbourCanada.AddSellableRessources("Tipiak")
	simu.env.AddHarbour(dockedHarbourCanada)

	//-------------------------------------------------------------
	// Mexique (Coordinates{23, 28})
	dockedHarbourMexique := NewHarbour("Mexique", Coordinates{23, 28})
	dockedHarbourMexique.AddAvailableRessources("Cocoa")
	dockedHarbourMexique.AddAvailableRessources("Cotton")
	dockedHarbourMexique.AddSellableRessources("Tobacco")
	dockedHarbourMexique.AddSellableRessources("Coffee")
	dockedHarbourMexique.AddSellableRessources("Spices")
	dockedHarbourMexique.AddSellableRessources("Sugar")
	simu.env.AddHarbour(dockedHarbourMexique)

	//-------------------------------------------------------------
	// Argentine (Coordinates{24, 39})
	dockedHarbourArgentine := NewHarbour("Argentine", Coordinates{24, 39})
	dockedHarbourArgentine.AddAvailableRessources("Spices")
	dockedHarbourArgentine.AddAvailableRessources("Weapons")
	dockedHarbourArgentine.AddAvailableRessources("Pearls")
	dockedHarbourArgentine.AddSellableRessources("Tipiak")
	dockedHarbourArgentine.AddSellableRessources("Rum")
	simu.env.AddHarbour(dockedHarbourArgentine)

	//-------------------------------------------------------------
	// Guadeloupe (Coordinates{25, 25})
	dockedHarbourGuadeloupe := NewHarbour("Guadeloupe", Coordinates{25, 25})
	dockedHarbourGuadeloupe.AddAvailableRessources("Sugar")
	dockedHarbourGuadeloupe.AddAvailableRessources("Spices")
	dockedHarbourGuadeloupe.AddSellableRessources("Rum")
	dockedHarbourGuadeloupe.AddSellableRessources("Tobacco")
	dockedHarbourGuadeloupe.AddSellableRessources("Cotton")
	simu.env.AddHarbour(dockedHarbourGuadeloupe)

	//-------------------------------------------------------------
	// Brésil (Coordinates{28, 31})
	dockedHarbourBresil := NewHarbour("Brésil", Coordinates{28, 31})
	dockedHarbourBresil.AddAvailableRessources("Coffee")
	dockedHarbourBresil.AddAvailableRessources("Sugar")
	dockedHarbourBresil.AddSellableRessources("Tobacco")
	dockedHarbourBresil.AddSellableRessources("Cotton")
	simu.env.AddHarbour(dockedHarbourBresil)

	//-------------------------------------------------------------
	// Groenland (Coordinates{37, 8})
	dockedHarbourGroenland := NewHarbour("Groenland", Coordinates{37, 8})
	dockedHarbourGroenland.AddAvailableRessources("Sugar")
	dockedHarbourGroenland.AddAvailableRessources("Tobacco")
	dockedHarbourGroenland.AddSellableRessources("Silver")
	dockedHarbourGroenland.AddSellableRessources("Weapons")
	simu.env.AddHarbour(dockedHarbourGroenland)

	//-------------------------------------------------------------
	// Compiègne / France (Coordinates{42, 22})
	dockedHarbourFrance := NewHarbour("Compiègne", Coordinates{42, 22})
	dockedHarbourFrance.AddAvailableRessources("Rum")
	dockedHarbourFrance.AddAvailableRessources("Cotton")
	dockedHarbourFrance.AddSellableRessources("Coffee")
	dockedHarbourFrance.AddSellableRessources("Sugar")
	simu.env.AddHarbour(dockedHarbourFrance)

	//-------------------------------------------------------------
	// Espagne (Coordinates{43, 25})
	dockedHarbourEspagne := NewHarbour("Espagne", Coordinates{43, 25})
	dockedHarbourEspagne.AddAvailableRessources("Rum")
	dockedHarbourEspagne.AddAvailableRessources("Tobacco")
	dockedHarbourEspagne.AddSellableRessources("Cotton")
	dockedHarbourEspagne.AddSellableRessources("Cocoa")
	simu.env.AddHarbour(dockedHarbourEspagne)

	//-------------------------------------------------------------
	// Angleterre (Coordinates{46, 18})
	dockedHarbourAngleterre := NewHarbour("Angleterre", Coordinates{46, 18})
	dockedHarbourAngleterre.AddAvailableRessources("Sugar")
	dockedHarbourAngleterre.AddAvailableRessources("Weapons")
	dockedHarbourAngleterre.AddSellableRessources("Rum")
	dockedHarbourAngleterre.AddSellableRessources("Cocoa")
	simu.env.AddHarbour(dockedHarbourAngleterre)

	//-------------------------------------------------------------
	// Congo (Coordinates{46, 35})
	dockedHarbourCongo := NewHarbour("Congo", Coordinates{46, 35})
	dockedHarbourCongo.AddAvailableRessources("Cotton")
	dockedHarbourCongo.AddAvailableRessources("Tobacco")
	dockedHarbourCongo.AddSellableRessources("Coffee")
	dockedHarbourCongo.AddSellableRessources("Cocoa")
	simu.env.AddHarbour(dockedHarbourCongo)

	//-------------------------------------------------------------
	// Algérie (Coordinates{50, 27})
	dockedHarbourAlgerie := NewHarbour("Algérie", Coordinates{50, 27})
	dockedHarbourAlgerie.AddAvailableRessources("Coffee")
	dockedHarbourAlgerie.AddAvailableRessources("Spices")
	dockedHarbourAlgerie.AddSellableRessources("Cotton")
	dockedHarbourAlgerie.AddSellableRessources("Tobacco")
	simu.env.AddHarbour(dockedHarbourAlgerie)

	//-------------------------------------------------------------
	// Afrique du Sud (Coordinates{49, 40})
	dockedHarbourAfriqueDuSud := NewHarbour("Afrique du Sud", Coordinates{49, 40})
	dockedHarbourAfriqueDuSud.AddAvailableRessources("Coffee")
	dockedHarbourAfriqueDuSud.AddAvailableRessources("Cotton")
	dockedHarbourAfriqueDuSud.AddSellableRessources("Tobacco")
	dockedHarbourAfriqueDuSud.AddSellableRessources("Rum")
	simu.env.AddHarbour(dockedHarbourAfriqueDuSud)

	//-------------------------------------------------------------
	// Russie (Coordinates{57, 14})
	dockedHarbourRussie := NewHarbour("Russie", Coordinates{57, 14})
	dockedHarbourRussie.AddAvailableRessources("Spices")
	dockedHarbourRussie.AddAvailableRessources("Tipiak")
	dockedHarbourRussie.AddAvailableRessources("Tobacco")
	dockedHarbourRussie.AddSellableRessources("Cocoa")
	dockedHarbourRussie.AddSellableRessources("Coffee")
	dockedHarbourArgentine.AddSellableRessources("Pearls")

	simu.env.AddHarbour(dockedHarbourRussie)

	//-------------------------------------------------------------
	// Inde (Coordinates{64, 30})
	dockedHarbourInde := NewHarbour("Inde", Coordinates{64, 30})
	dockedHarbourInde.AddAvailableRessources("Sugar")
	dockedHarbourInde.AddAvailableRessources("Ivory") // Exclusive
	dockedHarbourInde.AddAvailableRessources("Rum")
	dockedHarbourInde.AddSellableRessources("Silver")
	dockedHarbourInde.AddSellableRessources("Spices")
	simu.env.AddHarbour(dockedHarbourInde)

	//-------------------------------------------------------------
	// Chine (Coordinates{74, 25})
	dockedHarbourChine := NewHarbour("Chine", Coordinates{74, 25})
	dockedHarbourChine.AddAvailableRessources("Weapons")
	dockedHarbourChine.AddAvailableRessources("Tobacco")
	dockedHarbourChine.AddSellableRessources("Spices")
	dockedHarbourChine.AddSellableRessources("Sugar")
	dockedHarbourChine.AddSellableRessources("Silk")
	simu.env.AddHarbour(dockedHarbourChine)

	//-------------------------------------------------------------
	// Australie (Coordinates{75, 42})
	dockedHarbourAustralie := NewHarbour("Australie", Coordinates{75, 42})
	dockedHarbourAustralie.AddAvailableRessources("Silver")
	dockedHarbourAustralie.AddAvailableRessources("Cocoa")
	dockedHarbourAustralie.AddAvailableRessources("Weapons")
	dockedHarbourAustralie.AddSellableRessources("Spices")
	dockedHarbourAustralie.AddSellableRessources("Cotton")
	dockedHarbourAustralie.AddSellableRessources("Gold")
	simu.env.AddHarbour(dockedHarbourAustralie)

	//-------------------------------------------------------------
	// Japon (Coordinates{79, 27})
	dockedHarbourJapon := NewHarbour("Japon", Coordinates{79, 27})
	dockedHarbourJapon.AddAvailableRessources("Gold")
	dockedHarbourJapon.AddAvailableRessources("Silk")
	dockedHarbourJapon.AddSellableRessources("Coffee")
	dockedHarbourJapon.AddSellableRessources("Rum")
	dockedHarbourJapon.AddSellableRessources("Cocoa")
	simu.env.AddHarbour(dockedHarbourJapon)

	// Define nationalities
	nationalities := []constants.Nationality{
		constants.France,
		constants.England,
		constants.Spain,
	}

	// Specify how many merchant ships each nationality should have
	merchantCounts := map[constants.Nationality]int{
		constants.France:  50,
		constants.Spain:   50,
		constants.England: 50,
	}

	privateerCounts := map[constants.Nationality]int{
		constants.France:  10,
		constants.England: 10,
		constants.Spain:   10,
	}

	armyCounts := map[constants.Nationality]int{
		constants.France:  1,
		constants.England: 1,
		constants.Spain:   1,
	}

	numPirates := 50

	krakens := 2

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
	zoro := NewShipZoro("Zoro", NewRandomCoordinates(), nil)
	simu.env.AddAgent(zoro)

	rackham := NewShipRackham("Rackham", NewRandomCoordinates(), nil)
	simu.env.AddAgent(rackham)

	for i := 0; i < krakens; i++ {
		id := fmt.Sprintf("Kraken #%d", i)
		coor := NewRandomCoordinates()
		kraken := NewKraken(id, coor)
		simu.env.AddAgent(kraken)
	}

	// Create Pirate Ships (no nationality)
	for i := 0; i < numPirates; i++ {
		id := fmt.Sprintf("Pirate Ship #%d", i)
		coor := NewRandomCoordinates()
		ship := NewShipPirate(id, coor, nil)
		simu.env.AddAgent(ship)
	}

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
		time.Sleep(time.Second / (1 * time.Duration(simu.Parameters().TimeSpeed()))) // 1 step/s default and time speed
	}
}

func (simu *Simulation) Print() {

}

func (simu *Simulation) Log() {
	fmt.Print("starting")
}
