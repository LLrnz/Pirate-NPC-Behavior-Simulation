package simulation

import (
	"math"
	mathrand "math/rand"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
)

func pirateShipPercept(ship *Ship, env *Environment, step int) {

	kraken := env.KrakenNearby(ship)
	pirate := env.ClosestPirate(ship)
	merchant := env.ClosestMerchant(ship)
	privateer := env.ClosestPrivateer(ship)
	army := env.ClosestArmy(ship)

	if army != nil {
		//If nearby we set danger
		ship.SetDanger(army)
	} else {
		// retreive closest potential target between kraken pirate and privateer and merchant
		var distancesSlice = make([]int, 4)
		distancesSlice[0] = math.MaxInt // kraken to current ship distance
		distancesSlice[1] = math.MaxInt // pirate to current ship distance
		distancesSlice[2] = math.MaxInt // privateer to current ship distance
		distancesSlice[3] = math.MaxInt // merchant to current ship distance

		if kraken != nil {
			distancesSlice[0] = Heuristic((*kraken).Coordinates(), ship.coordinates)
		}
		if pirate != nil {
			distancesSlice[1] = Heuristic((pirate).Coordinates(), ship.coordinates)
		}
		if privateer != nil {
			distancesSlice[2] = Heuristic((privateer).Coordinates(), ship.coordinates)
		}
		if merchant != nil {
			distancesSlice[3] = Heuristic((merchant).Coordinates(), ship.coordinates)
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
			ship.SetTarget(*kraken)
		} else if indiceMin == 1 && pirate != nil {
			ship.SetTarget(pirate)
		} else if indiceMin == 2 && privateer != nil {
			ship.SetTarget(privateer)
		} else if indiceMin == 3 && merchant != nil {
			ship.SetTarget(merchant)
		} else {
		}
	}
}

func pirateShipDeliberate(ship *Ship, env *Environment, step int) {

	// if docked and not full life, fix or sell ressources
	if ship.Docked() != nil {

		// If not max health, repair ship
		if ship.Health() != constants.ShipConfigs["Pirate"].Health && ship.Money() >= 1 {
			ship.SetAction("repair")
			return
		}

		// If ressources in ship that can be sold at harbour, sell
		if env.HasRessourcesToSellAtHarbour(ship) && ship.Quantity() >= 1 {
			ship.SetAction("sell")
			return
		}

		// If ship is in danger wait
		if ship.Danger() != nil {
			ship.SetAction("wait")
			return
		}

	}

	// if docked, set itinerary to another harbour
	if ship.Docked() != nil {
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

			// Set ship itinerary to the closest harbor where resources can be sold
			path := AStar(ship.Coordinates(), closestHarbour.Coordinates())
			if path != nil {
				ship.SetItinerary(path)
			}

		} else {
			harbours := env.Harbours()
			numberHarbour := mathrand.Intn(len(harbours) - 1)
			destinationHarbour := harbours[numberHarbour]

			// Set ship itinerary to random harbor
			path := AStar(ship.Coordinates(), destinationHarbour.Coordinates())
			if path != nil {
				ship.SetItinerary(path)
			}

		}

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

		ship.SetAction("follow")
		return
	}

	// if ship has a target
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target

			ship.SetAction("attack")
		} else { // else, follow target
			// Pathfinding to the target
			path := AStar(ship.Coordinates(), (ship.Target()).Coordinates())
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
				ship.SetAction("follow")
			}
		}
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

	if ship.Danger() != nil {
		ship.SetAction("flee")
	} else {
		ship.SetAction("follow")
	}
}

func pirateShipAct(ship *Ship, env *Environment, step int) {
	switch ship.Action() {
	// case "flee":
	// 	// Continue itinerary to flee # see if pirate can attack kraken or not
	// 	ship.FollowItinerary()
	case "flee":
		// Continue itinerary to flee
		ship.FollowItinerary()
	case "repair":
		ship.SetHealth(constants.ShipConfigs["Pirate"].Health)
	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		// Remove the target from the environment
		env.AttackAgent(ship.Target(), ship) // need to stell ressources from the target
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

	default:
		// Do nothing
	}
}

func pirateShipLevelUp(ship *Ship, env *Environment, step int) {
	//
	// pirateShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Merchant ship -------------------

func merchantShipPercept(ship *Ship, env *Environment, step int) {
	//Verify if kraken is nearby within Fov
	kraken := env.KrakenNearby(ship)

	if kraken != nil {
		//If nearby we set danger
		ship.SetDanger(*kraken)
	} else {
		//Else danger is nil
		ship.SetDanger(nil)
	}

}

func merchantShipDeliberate(ship *Ship, env *Environment, step int) {

	// If kraken within fov go 3 cases away from it or as far as possible max 3 cases
	if ship.Danger() != nil && ship.Docked() == nil {
		harbour := env.ClosestHarbour(ship)
		path := AStar(ship.Coordinates(), harbour.Coordinates())
		// If path found set itinerary and follow
		if path != nil {
			ship.SetItinerary(path)
			ship.SetAction("flee")
		} else {
			ship.SetAction("wait")
		}
		return
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
	print(ship.Action(), "\n")
	switch ship.Action() {
	case "flee":
		// Continue itinerary to flee
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
	ship.danger = nil
	ship.action = ""

}

func merchantShipLevelUp(ship *Ship, env *Environment, step int) {
	// merchantShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Privateer ship -------------------

func privateerShipPercept(ship *Ship, env *Environment, step int) {
	kraken := env.KrakenNearby(ship)
	pirate := env.ClosestPirate(ship)
	merchant := env.ClosestMerchantNationality(ship)
	privateer := env.ClosestPrivateerNationality(ship)
	army := env.ClosestArmyNationality(ship)

	// retreive closest potential target between kraken pirate and privateer and merchant
	var distancesSlice = make([]int, 5)
	distancesSlice[0] = math.MaxInt // kraken to current ship distance
	distancesSlice[1] = math.MaxInt // pirate to current ship distance
	distancesSlice[2] = math.MaxInt // privateer to current ship distance
	distancesSlice[3] = math.MaxInt // merchant to current ship distance

	if kraken != nil {
		distancesSlice[0] = Heuristic((*kraken).Coordinates(), ship.coordinates)
	}
	if pirate != nil {
		distancesSlice[1] = Heuristic((pirate).Coordinates(), ship.coordinates)
	}
	if privateer != nil {
		distancesSlice[2] = Heuristic((privateer).Coordinates(), ship.coordinates)
	}
	if merchant != nil {
		distancesSlice[3] = Heuristic((merchant).Coordinates(), ship.coordinates)
	}
	if army != nil {
		distancesSlice[4] = Heuristic((army).Coordinates(), ship.coordinates)
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
		ship.SetTarget(*kraken)
	} else if indiceMin == 1 && pirate != nil {
		ship.SetTarget(pirate)
	} else if indiceMin == 2 && privateer != nil {
		ship.SetTarget(privateer)
	} else if indiceMin == 3 && merchant != nil {
		ship.SetTarget(merchant)
	} else if indiceMin == 4 && army != nil {
		ship.SetTarget(army)
	}
}

func privateerShipDeliberate(ship *Ship, env *Environment, step int) {

	// if docked and not full life, fix or sell ressources
	if ship.Docked() != nil {

		// If not max health, repair ship
		if ship.Health() != constants.ShipConfigs["Privateer"].Health && ship.Money() >= 1 {
			ship.SetAction("repair")
			return
		}

		// If ressources in ship that can be sold at harbour, sell
		if env.HasRessourcesToSellAtHarbour(ship) && ship.Quantity() >= 1 {
			ship.SetAction("sell")
			return
		}

	}

	// if docked, set itinerary to another harbour
	if ship.Docked() != nil {
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

			// Set ship itinerary to the closest harbor where resources can be sold
			path := AStar(ship.Coordinates(), closestHarbour.Coordinates())
			if path != nil {
				ship.SetItinerary(path)
			}

		} else {
			harbours := env.Harbours()
			numberHarbour := mathrand.Intn(len(harbours) - 1)
			destinationHarbour := harbours[numberHarbour]

			// Set ship itinerary to random harbor
			path := AStar(ship.Coordinates(), destinationHarbour.Coordinates())
			if path != nil {
				ship.SetItinerary(path)
			}

		}

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

		ship.SetAction("follow")
		return
	}

	// if ship has a target
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target

			ship.SetAction("attack")
		} else { // else, follow target
			// Pathfinding to the target
			path := AStar(ship.Coordinates(), (ship.Target()).Coordinates())
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
				ship.SetAction("follow")
			}
		}
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

	ship.SetAction("follow")

}

func privateerShipAct(ship *Ship, env *Environment, step int) {
	switch ship.Action() {
	case "repair":
		ship.SetHealth(constants.ShipConfigs["Privateer"].Health)
	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		// Remove the target from the environment
		env.AttackAgent(ship.Target(), ship)
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

func privateerShipLevelUp(ship *Ship, env *Environment, step int) {
	// privateerShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Army ship -------------------

func armyShipPercept(ship *Ship, env *Environment, step int) {

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
		distancesSlice[1] = Heuristic((pirate).Coordinates(), ship.coordinates)
	}
	if privateer != nil {
		distancesSlice[2] = Heuristic((privateer).Coordinates(), ship.coordinates)
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
		ship.SetTarget(*kraken)
	} else if indiceMin == 1 && pirate != nil {
		ship.SetTarget(pirate)
	} else if indiceMin == 2 && privateer != nil {
		ship.SetTarget(privateer)
	}
}

func armyShipDeliberate(ship *Ship, env *Environment, step int) {

	// actions in priority order

	// if docked and not full life, fix
	if ship.Docked() != nil && ship.Health() < constants.ShipConfigs["Army"].Health {
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

		ship.SetAction("follow")
		return
	}

	// if ship has a target
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target

			ship.SetAction("attack")
		} else { // else, follow target
			// Pathfinding to the target
			path := AStar(ship.Coordinates(), (ship.Target()).Coordinates())
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
				ship.SetAction("follow")
			}
		}
		return
	}

	// if ship has no target, follow current itinerary
	if len(ship.Itinerary()) != 0 || ship.Itinerary() != nil {
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

	ship.SetAction("follow")
}

func armyShipAct(ship *Ship, env *Environment, step int) {
	// Army ship can attack Pirate ship,  Filibusteer ship, Privateer ship

	switch ship.Action() {
	case "repair":
		ship.SetHealth(constants.ShipConfigs["Army"].Health)
	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		env.AttackAgent(ship.Target(), ship)
	case "follow":
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

	// Update current shipState and direction for sprite
	ship.CurrentShipState()
	ship.CurrentDirection()
}

func armyShipLevelUp(ship *Ship, env *Environment, step int) {
	// Army ship can't level up i think

}

// ------------------- Zoro ship -------------------

func zoroShipPercept(ship *Ship, env *Environment, step int) {

	// Zoro percepts a target within its fov
	// Zoro's target is any ship/kraken who's closest to him

	kraken := env.KrakenNearby(ship)
	targetShip := env.ClosestShipFromShip(ship)

	var distancesSlice = make([]int, 2)
	distancesSlice[0] = math.MaxInt // Kraken
	distancesSlice[1] = math.MaxInt // targetShip

	if kraken != nil {
		distancesSlice[0] = Heuristic((*kraken).Coordinates(), ship.coordinates)
	}
	if targetShip != nil {
		distancesSlice[0] = Heuristic((targetShip).Coordinates(), ship.coordinates)
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
		ship.SetTarget(*kraken)
	} else if indiceMin == 1 && targetShip != nil {
		ship.SetTarget(targetShip)
	} else {
	}

}

func zoroShipDeliberate(ship *Ship, env *Environment, step int) {
	// if docked and not full life, fix
	if ship.Docked() != nil && ship.Health() < constants.ShipConfigs["Zoro"].Health {
		ship.SetAction("repair")
		return
	}

	// since Zoro is lost at all times, he can't decide to go to the closest harbour when his life is low

	// if docked, set itinerary to somewhere else (random)
	if ship.Docked() != nil {
		destination := NewRandomCoordinates()
		path := AStar(ship.Coordinates(), destination)
		if path != nil {
			ship.SetItinerary(path)
		}
		ship.SetAction("follow")
		return
	}

	// If Zoro has a target, attack/follow it
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target
			ship.SetAction("attack")
		} else { // else, follow target // should never happen since Zoro should only detect agt on the same coordinates as him
			// Pathfinding to the target
			path := AStar(ship.Coordinates(), (ship.Target()).Coordinates())
			// if path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
				ship.SetAction("follow")
			}
		}
		return
	}

	// If Zoro has no target, follow current itinerary
	if ship.Itinerary() != nil {
		ship.SetAction("follow")
		return
	}

	// If Zoro has no target and no current itinerary, compute new itinerary
	// get random coordinates
	destination := NewRandomCoordinates()

	// Set ship itinerary to chosen destination
	path := AStar(ship.Coordinates(), destination)
	if path != nil {
		ship.SetItinerary(path)
	}

	ship.SetAction("follow")
}

func zoroShipAct(ship *Ship, env *Environment, step int) {
	//LostZoro(ship, env)

	switch ship.Action() {
	case "repair":
		ship.SetHealth(constants.ShipConfigs["Zoro"].Health)
	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		env.AttackAgent(ship.Target(), ship)
	case "follow":
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

func zoroShipLevelUp(ship *Ship, env *Environment, step int) {
	// no level up for zoro ship, can't access to the port
}

// ------------------- Rackham ship -------------------

func rackhamShipPercept(ship *Ship, env *Environment, step int) {
	kraken := env.KrakenNearby(ship)
	pirate := env.ClosestPirate(ship)
	merchant := env.ClosestMerchant(ship)
	privateer := env.ClosestPrivateer(ship)
	army := env.ClosestArmy(ship)

	// retreive closest potential target between kraken pirate and privateer and merchant
	var distancesSlice = make([]int, 5)
	distancesSlice[0] = math.MaxInt // kraken to current ship distance
	distancesSlice[1] = math.MaxInt // pirate to current ship distance
	distancesSlice[2] = math.MaxInt // privateer to current ship distance
	distancesSlice[3] = math.MaxInt // merchant to current ship distance
	distancesSlice[4] = math.MaxInt // merchant to current ship distance

	if kraken != nil {
		distancesSlice[0] = Heuristic((*kraken).Coordinates(), ship.coordinates)
	}
	if pirate != nil {
		distancesSlice[1] = Heuristic((pirate).Coordinates(), ship.coordinates)
	}
	if privateer != nil { // target only other nationality privateer
		distancesSlice[2] = Heuristic((privateer).Coordinates(), ship.coordinates)
	}
	if merchant != nil { // target all nationality merchant
		distancesSlice[3] = Heuristic((merchant).Coordinates(), ship.coordinates)
	}

	if army != nil { // target all nationality army
		distancesSlice[4] = Heuristic((army).Coordinates(), ship.coordinates)
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
		ship.SetTarget(*kraken)
	} else if indiceMin == 1 && pirate != nil {
		ship.SetTarget(pirate)
	} else if indiceMin == 2 && privateer != nil {
		ship.SetTarget(privateer)
	} else if indiceMin == 3 && merchant != nil {
		ship.SetTarget(merchant)
	} else if indiceMin == 4 && army != nil {
		ship.SetTarget(army)
	} else {
	}
}

func rackhamShipDeliberate(ship *Ship, env *Environment, step int) {
	// if docked and not full life, fix or sell ressources
	if ship.Docked() != nil {

		// If not max health, repair ship
		if ship.Health() != constants.ShipConfigs["Rackham"].Health && ship.Money() >= 1 {
			ship.SetAction("repair")
			return
		}

		// If ressources in ship that can be sold at harbour, sell
		if env.HasRessourcesToSellAtHarbour(ship) && ship.Quantity() >= 1 {
			ship.SetAction("sell")
			return
		}

	}

	// if docked, set itinerary to another harbour
	if ship.Docked() != nil {
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

			// Set ship itinerary to the closest harbor where resources can be sold
			path := AStar(ship.Coordinates(), closestHarbour.Coordinates())
			if path != nil {
				ship.SetItinerary(path)
			}

		} else {
			harbours := env.Harbours()
			numberHarbour := mathrand.Intn(len(harbours) - 1)
			destinationHarbour := harbours[numberHarbour]

			// Set ship itinerary to random harbor
			path := AStar(ship.Coordinates(), destinationHarbour.Coordinates())
			if path != nil {
				ship.SetItinerary(path)
			}

		}

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

		ship.SetAction("follow")
		return
	}

	// if ship has a target
	if ship.Target() != nil {
		distance := Heuristic(ship.Coordinates(), (ship.Target()).Coordinates())

		if distance <= ship.Speed() { // if ship is next to a target, attack target

			ship.SetAction("attack")
		} else { // else, follow target
			// Pathfinding to the target
			path := AStar(ship.Coordinates(), (ship.Target()).Coordinates())
			// If path found set itinerary and follow
			if path != nil {
				ship.SetItinerary(path)
				ship.SetAction("follow")
			}
		}
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

	if ship.Danger() != nil {
		ship.SetAction("flee")
	} else {
		ship.SetAction("follow")
	}
}

func rackhamShipAct(ship *Ship, env *Environment, step int) {
	switch ship.Action() {

	case "flee":
		// Continue itinerary to flee
		ship.FollowItinerary()
	case "repair":
		ship.SetHealth(constants.ShipConfigs["Army"].Health)
	case "attack":
		// Go to target and attack
		ship.FollowItinerary()
		// Remove the target from the environment
		env.AttackAgent(ship.Target(), ship) // need to stell ressources from the target
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

	default:
		// Do nothing
	}
}

func rackhamShipLevelUp(ship *Ship, env *Environment, step int) {
}
