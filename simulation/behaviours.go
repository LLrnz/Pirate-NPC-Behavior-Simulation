package simulation

import (
	"fmt"
	"math"

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
		if env.HasRessourcesToSellAtHarbour(ship) {
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

			for ressource, available := range resources {
				if !available {
					continue
				}

				if ship.Ressources()[ressource] > 0 {
					// Calculate distance to the harbor
					distance := float32(Heuristic(ship.Coordinates(), harbour.Coordinates()))

					if distance < shortestDistance {
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
		availableCapacity := ship.Quantity()
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
	//fmt.Println("Army ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

	// army's targets are pirate ships, privateer ships (from different nationality) and kraken
	// retreive potential target from environment

	//kraken := env.KrakenNearby(ship)
	//pirate := env.ClosestPirate(ship)
	//privateer := env.ClosestPrivateer(ship)
	//target := closest agt from the previous

	/*if target != nil {
		fmt.Println((*target).Id())
		fmt.Println((*target).Coordinates())
	} else {
		fmt.Println("Army ship doesn't detect any target")
	}
	*/

	//ship.SetTarget(target)
}

func armyShipDeliberate(ship *Ship, env *Environment, step int) {
	//fmt.Println("Army ship deliberate...")

	// actions in priority order

	// if docked and not full life, fix
	if ship.Docked() != nil && ship.Health() < maxHealth {
		ship.SetAction("repair")
		return
	}

	// if docked, set itinerary to another harbour
	if ship.Docked() != nil {
		ship.SetAction("toAnyHarbour")
		return
	}

	// if life under 20, set itinerary to closest harbour
	if ship.Health() <= 20 {
		ship.SetAction("toClosestHarbour")
		return
	}

	// if ship has a target
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (*ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target
			ship.SetAction("attack")
		} else { // else, follow target
			ship.SetAction("follow")
		}
		return
	}
}

func armyShipAct(ship *Ship, env *Environment, step int) {
	//fmt.Println("Army ship act...")
	// Army ship can attack Pirate ship,  Filibusteer ship, Privateer ship

	switch ship.Action() {
	case "repair":
		ship.SetHealth(maxHealth)
	case "toAnyHarbour":
		// set itinerary to another harbour (except current one)
		// retreive any harbour (except current one) from environment
		var anyHarbour *Harbour

		distance := Heuristic(ship.Coordinates(), (*anyHarbour).Coordinates())

		if distance <= ship.Speed() { // if ship is next to harbour, docked

			// TODO : check if there's a free spot in the harbour -> if not, wait for a free spot ?
			ship.SetDocked(anyHarbour)
		} else {
			path := AStar(constants.MapGrid, ship.Coordinates(), (*ship.Target()).Coordinates())

			fmt.Println(path)
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
			}
		}

		ship.FollowItinerary()

	case "toClosestHarbour":
		// retreive closest harbour from environment
		var closestHarbour *Harbour

		distance := Heuristic(ship.Coordinates(), (*closestHarbour).Coordinates())

		if distance <= ship.Speed() { // if ship is next to harbour, docked
			ship.SetDocked(closestHarbour)
		} else {
			path := AStar(constants.MapGrid, ship.Coordinates(), (*ship.Target()).Coordinates())

			fmt.Println(path)
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
			}
		}

		ship.FollowItinerary()

	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		// Remove the target from the environment
		env.AttackAgent(*ship.Target(), ship)

	case "follow":
		// Follow target
		ship.FollowItinerary()

	default:
		ship.FollowItinerary()
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
