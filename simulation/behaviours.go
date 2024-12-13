package simulation

import (
	"fmt"
	"math"
	mathrand "math/rand"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
)

func pirateShipPercept(ship *Ship, env *Environment, step int) {
	//fmt.Println("Pirate ship percept...")

	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func pirateShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Pirate ship deliberate...")
	// ClosestAgentForPirate(ship) //regarde si elle est good pour choisir la cible
}

func pirateShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Pirate ship act...")
	// Pirate ship can attack Merchant ship, Filibusteer ship, Privateer ship, Army ship
}

func pirateShipLevelUp(ship *Ship, env *Environment, step int) {
	//
	//fmt.Println("Pirate ship level up...")
	// pirateShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Merchant ship -------------------

func merchantShipPercept(ship *Ship, env *Environment, step int) {
	//Verify if kraken is nearby within Fov
	kraken := env.KrakenNearby(ship)

	if kraken != nil {
		//If nearby we set danger
		ship.SetDanger(kraken)
	} else {
		//Else danger is nil
		ship.SetDanger(nil)
	}

}

func merchantShipDeliberate(ship *Ship, env *Environment, step int) {

	// If kraken within fov go 3 cases away from it or as far as possible max 3 cases
	if ship.Danger() != nil && ship.Docked() == nil {
		// Get ship and kraken coordinates
		shipCoords := ship.Coordinates()
		krakenCoords := (*ship.Danger()).Coordinates()

		// Calculate the direction where the ship should go
		dx := shipCoords.X() - krakenCoords.X()
		dy := shipCoords.Y() - krakenCoords.Y()

		magnitude := math.Sqrt(float64(dx*dx + dy*dy))
		dirX := int(math.Round(float64(dx) / magnitude))
		dirY := int(math.Round(float64(dy) / magnitude))

		// Attempt to move 3 cases away in the opposite direction
		for cases := 3; cases > 0; cases-- {
			targetX := shipCoords.X() + dirX*cases
			targetY := shipCoords.Y() + dirY*cases

			// Check if the target coordinate is within bounds and walkable
			if targetX >= 0 && targetX < len(constants.MapGrid[0]) &&
				targetY >= 0 && targetY < len(constants.MapGrid) &&
				constants.MapGrid[targetY][targetX] == 0 {

				// Pathfinding to the target
				path := AStar(ship.Coordinates(), Coordinates{x: targetX, y: targetY})
				// If path found set itinerary and follow
				if path != nil {
					ship.SetItinerary(path)
					ship.SetAction("flee")
				} else {
					ship.SetAction("wait")
				}
				return
			}
		}
	}

	// Else if path to follow not finished, follow it
	if len(ship.Itinerary()) != 0 {
		ship.SetAction("follow")
		return
	}

	// Else if docked at a harbour, either transaction if ressources or fix ship or get to other port
	if ship.Docked() != nil {

		// If not max health, repair ship
		if ship.Health() != constants.ShipConfigs["Merchant"].Health && ship.Money() >= 1 {
			ship.SetAction("repair")
			return
		}

		// If ressources in ship that can be sold at harbour, sell
		if env.HasRessourcesToSellAtHarbour(ship) && ship.Quantity() >= 1 {
			ship.SetAction("sell")
			return
		}

		// If ship is not full take ressources
		if ship.Quantity() == 0 {
			ship.SetAction("buy")
			return
		}

		if ship.Danger() != nil {
			ship.SetAction("wait")
			return
		}

	}

	// If no itinerary and nothing to do at harbour, recalculate itinerary
	totalResources := ship.Quantity()

	if totalResources > 0 {
		// Go to the closest harbor where resources can be sold
		harbours := env.Harbours()
		closestHarbour := harbours[0]
		shortestDistance := float32(-1)

		for _, harbour := range harbours {
			resources := harbour.SellableRessources()

			if len(resources) == 0 {
				continue
			}

			for ressource, available := range resources {
				if !available {
					continue
				}

				if ship.Ressources()[ressource] > 0 {
					// Calculate distance to the harbor
					distance := float32(Heuristic(ship.Coordinates(), harbour.Coordinates()))
					if distance < shortestDistance || shortestDistance == -1 {
						shortestDistance = distance
						closestHarbour = harbour
					}

				}
			}
		}

		// Set ship itinerary to the closest harbor
		path := AStar(ship.Coordinates(), closestHarbour.Coordinates())
		if path != nil {
			ship.SetItinerary(path)
		}

	} else {
		harbours := env.Harbours()
		bestHarbour := harbours[0]
		bestScore := float32(-1)

		for _, harbour := range harbours {
			// Calculate distance heuristic
			distance := Heuristic(ship.Coordinates(), harbour.Coordinates())
			distanceScore := 1 / (1 + float32(distance))

			// Evaluate the best resource at the harbor
			resources := harbour.AvailableRessources()
			highestResourcePrice := float32(0)

			for ressource, available := range resources {
				if !available {
					continue
				}

				ressourceObj, ok := GetRessource(ressource)
				if ok {
					// Calculate the market price of the resource
					price := env.Market().CalculateRessourcePrice(step, ressourceObj)

					// Highest price resource
					if price > highestResourcePrice {
						highestResourcePrice = price
					}
				}
			}

			// Combine scores using the best resource price
			totalScore := distanceScore * highestResourcePrice

			// Update best harbor
			if totalScore > bestScore {
				bestScore = totalScore
				bestHarbour = harbour
			}
		}

		// After selecting the best harbor, set the itinerary
		path := AStar(ship.Coordinates(), bestHarbour.Coordinates())
		if path != nil {
			ship.SetItinerary(path)
		}

		ship.SetAction("follow")
		return
	}

	ship.SetAction("follow")

}

func merchantShipAct(ship *Ship, env *Environment, step int) {
	switch ship.Action() {
	case "flee":
		// Continue itinerary to flee
		ship.FollowItinerary()
	case "follow":
		// Continue itinerary
		ship.FollowItinerary()

		if ship.Docked() != nil {
			env.UndockShip(ship)
		}

		if len(ship.Itinerary()) == 0 {
			harbour, ok := env.IsHarbour(ship)
			if ok {
				env.DockShip(ship, harbour)
			}
		}
	case "repair":
		// Repair ship for money
		ship.Repair()
	case "sell":
		for ressource, available := range ship.Docked().SellableRessources() {
			// Check if the resource is sellable
			if !available {
				continue
			}

			// Check if the ship has this resource and the quantity is greater than 0
			quantity := ship.Ressources()[ressource]
			if quantity <= 0 {
				continue
			}

			//Get the Ressource Object of the ressource
			ressourceObj, ok := GetRessource(ressource)
			if ok {
				// Sell
				env.Market().Sell(ship, ressourceObj, step)
			}
		}
	case "buy":
		// Evaluate resources at the harbor
		resources := ship.Docked().AvailableRessources()
		var mostExpensiveRessource Ressource
		highestPrice := float32(0)

		// Identify the most expensive resource
		for ressource, available := range resources {
			if !available {
				continue
			}

			ressourceObj, ok := GetRessource(ressource)
			if ok {
				// Calculate the market price of the resource
				price := env.Market().CalculateRessourcePrice(step, ressourceObj)
				if price > highestPrice {
					highestPrice = price
					mostExpensiveRessource = ressourceObj
				}
			}
		}

		// Calculate the amount to load based on the ship capacity
		maxCapacity := constants.ShipConfigs["Merchant"].Capacity
		availableCapacity := maxCapacity - ship.Quantity()
		quantityToLoad := min(maxCapacity, availableCapacity)

		// Buy ressource for the ship
		ship.AddRessources(mostExpensiveRessource.name, quantityToLoad)

	case "wait":
		// Do nothing
	default:
		// Do nothing
	}

	ship.target = nil
	ship.action = ""

}

func merchantShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("Merchant ship level up...")
	// merchantShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Privateer ship -------------------

func privateerShipPercept(ship *Ship, env *Environment, step int) {
	//fmt.Println("Privateer ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func privateerShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Privateer ship deliberate...")
}

func privateerShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Privateer ship act...")
	// Privateer ship can attack Pirate ship, Merchant ship, Filibusteer ship, Army ship
}

func privateerShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("Privateer ship level up...")
	// privateerShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Army ship -------------------

func armyShipPercept(ship *Ship, env *Environment, step int) {
	fmt.Println("Army ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

	// army's targets are pirate ships, privateer ships (from different nationality) and kraken
	// retreive potential target from environment

	kraken := env.KrakenNearby(ship)
	pirate := env.ClosestPirate(ship)
	privateer := env.ClosestPrivateerNationality(ship)

	// retreive closest potential target between kraken pirate and privateer
	var distancesSlice = make([]int, 3)
	distancesSlice[0] = math.MaxInt // kraken to current ship distance
	distancesSlice[1] = math.MaxInt // pirate to current ship distance
	distancesSlice[2] = math.MaxInt // privateer to current ship distance

	if kraken != nil {
		distancesSlice[0] = Heuristic((*kraken).Coordinates(), ship.coordinates)
	}
	if pirate != nil {
		distancesSlice[1] = Heuristic((*pirate).Coordinates(), ship.coordinates)
	}
	if privateer != nil {
		distancesSlice[2] = Heuristic((*privateer).Coordinates(), ship.coordinates)
	}

	distanceMin := distancesSlice[0]
	indiceMin := 0

	for i, d := range distancesSlice {
		if d < distanceMin {
			indiceMin = i
			distanceMin = d
		}
	}

	if indiceMin == 0 && kraken != nil {
		fmt.Println((*kraken).Id())
		fmt.Println((*kraken).Coordinates())
		ship.SetTarget(kraken)
	} else if indiceMin == 1 && pirate != nil {
		fmt.Println((*pirate).Id())
		fmt.Println((*pirate).Coordinates())
		ship.SetTarget(pirate)
	} else if indiceMin == 2 && privateer != nil {
		fmt.Println((*privateer).Id())
		fmt.Println((*privateer).Coordinates())
		ship.SetTarget(privateer)
	} else {
		fmt.Println("Army ship doesn't detect any target")
	}
}

func armyShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Army ship deliberate...")

	// actions in priority order

	// if docked and not full life, fix
	if ship.Docked() != nil && ship.Health() < maxHealth {
		fmt.Println("Army ship deliberate - repair")
		ship.SetAction("repair")
		return
	}

	// if docked, set itinerary to another harbour
	if ship.Docked() != nil {
		// retreive harbours from environment
		harbours := env.Harbours()

		// remove current harbour from the list
		/*for i, harbour := range harbours {
			if harbour.Coordinates() == ship.Docked().Coordinates() {
				harbours = append(harbours[:i], harbours[i+1:]...)
				break
			}
		}*/

		// get random harbour from the list
		numberHarbour := mathrand.Intn(len(harbours) - 1)
		destinationHarbour := harbours[numberHarbour]

		// Set ship itinerary to chosen harbor
		path := AStar(ship.Coordinates(), destinationHarbour.Coordinates())
		if path != nil {
			ship.SetItinerary(path)
		}

		fmt.Println("Army ship deliberate - docked, set itinerary to another harbour - follow")
		ship.SetAction("follow")

		return
	}

	// if life under 20, set itinerary to closest harbour
	if ship.Health() <= 20 {

		// retreive harbours from environment
		harbours := env.Harbours()
		closestHarbour := harbours[0]
		shortestDistance := float32(-1)

		for _, harbour := range harbours {
			// Calculate distance to harbor
			distance := float32(Heuristic(ship.Coordinates(), harbour.Coordinates()))

			if distance < shortestDistance {
				shortestDistance = distance
				closestHarbour = harbour
			}

		}

		// Set ship itinerary to the closest harbor
		path := AStar(ship.Coordinates(), closestHarbour.Coordinates())
		if path != nil {
			ship.SetItinerary(path)
		}

		fmt.Println("Army ship deliberate - life under 20, set itinerary to closest harbour - follow")
		ship.SetAction("follow")
		return
	}

	// if ship has a target
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (*ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target

			fmt.Println("Army ship deliberate - target - attack")
			ship.SetAction("attack")
		} else { // else, follow target
			// Pathfinding to the target
			path := AStar(ship.Coordinates(), (*ship.Target()).Coordinates())
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
				fmt.Println("Army ship deliberate - target - follow")
				ship.SetAction("follow")
			}
		}
		return
	}

	// if ship has no target, follow current itinerary
	if ship.Itinerary() != nil {
		fmt.Println("Army ship deliberate - no target, follow current itinerary - follow")
		ship.SetAction("follow")
		return
	}

	// ship itinerary is null
	// retreive harbours from environment
	harbours := env.Harbours()

	// get random harbour from the list
	numberHarbour := mathrand.Intn(len(harbours) - 1)
	destinationHarbour := harbours[numberHarbour]

	// Set ship itinerary to chosen harbor
	path := AStar(ship.Coordinates(), destinationHarbour.Coordinates())
	if path != nil {
		ship.SetItinerary(path)
	}

	fmt.Println("Army ship deliberate - no target, no itinerary - follow")
	ship.SetAction("follow")
}

func armyShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Army ship act...")
	// Army ship can attack Pirate ship,  Filibusteer ship, Privateer ship

	switch ship.Action() {
	case "repair":
		ship.SetHealth(maxHealth)
	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		// Remove the target from the environment
		env.AttackAgent(*ship.Target(), ship)
	case "follow":
		// Follow target
		ship.FollowItinerary()

		if ship.Docked() != nil {
			env.UndockShip(ship)
		}

		if len(ship.Itinerary()) == 0 {
			harbour, ok := env.IsHarbour(ship)
			if ok {
				env.DockShip(ship, harbour)
			}
		}
	default:
		// Do nothing
	}
}

func armyShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("Army ship level up is not possible...")
	// Army ship can't level up i think

}

// ------------------- Filibusteer ship -------------------

func filibusteerShipPercept(ship *Ship, env *Environment, step int) {
	//fmt.Println("Filibusteer ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func filibusteerShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Filibusteer ship deliberate...")
}

func filibusteerShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Filibusteer ship act...")
}

func filibusteerShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("Filibusteer ship level up...")
	// filibusteerpShipLevelUp(ship *Ship, env *Environment)
}

// ------------------- Pirate hunter ship -------------------

func piratehunterShipPercept(ship *Ship, env *Environment, step int) {
	//fmt.Println("PirateHunter ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func piratehunterShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("PirateHunter ship deliberate...")
}

func piratehunterShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("PirateHunter ship act...")
	// PirateHunter ship can attack Pirate ship, Filibusteer ship, Privateer ship

}

func piratehunterShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("PirateHunter ship level up...")
	// piratehunterShipLevelUp(ship *Ship, env *Environment)
}

// ------------------- Zoro ship -------------------

func zoroShipPercept(ship *Ship, env *Environment, step int) {
	//fmt.Println("Zoro ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby
}

func zoroShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Zoro ship deliberate...")
}

func zoroShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Zoro ship act...")
	//LostZoro(ship, env)

}

func zoroShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("Zoro ship level up is not possible...")
	// no level up for zoro ship, can't access to the port
}

// ------------------- Rackham ship -------------------

func rackhamShipPercept(ship *Ship, env *Environment, step int) {

	//fmt.Println("Rackham ship percept...")
}

func rackhamShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Rackham ship deliberate...")
}

func rackhamShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Rackham ship act...")

}

func rackhamShipLevelUp(ship *Ship, env *Environment, step int) {
	//fmt.Println("Rackham ship level up is not possible...")
}
