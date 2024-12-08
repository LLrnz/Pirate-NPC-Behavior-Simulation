package simulation

func pirateShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("Pirate ship percept...")

	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func pirateShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("Pirate ship deliberate...")
	// ClosestAgentForPirate(ship) //regarde si elle est good pour choisir la cible
}

func pirateShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Pirate ship act...")
	// Pirate ship can attack Merchant ship, Filibusteer ship, Privateer ship, Army ship
}

func pirateShipLevelUp(ship *Ship, env *Environment) {
	//
	//fmt.Println("Pirate ship level up...")
	// pirateShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Merchant ship -------------------

func merchantShipPercept(ship *Ship, env *Environment) {
	//Verify if kraken is nearby within Fov
	kraken := env.KrakenNearby(ship)

	if kraken != nil {
		//If nearby we set danger
		ship.SetDanger(&kraken)
	} else {
		//Else danger is nil
		ship.SetDanger(nil)
	}

}

func merchantShipDeliberate(ship *Ship, env *Environment) {
	// If kraken nearby, get away as priority
	if ship.danger != nil {
		ship.SetAction("flee")
		return
	}

	// Else if path to follow not finished, follow it
	if len(ship.itinerary) != 0 {
		ship.SetAction("follow")
		return
	}

	// Else if docked at a harbour, either transaction if ressources or fix ship or get to other port
	if ship.Docked() != nil {

		// If not max health, repair ship
		if ship.Health() != maxHealth {
			ship.SetAction("repair")
			return
		}

		// If ressources in ship that can be sold at harbour, sell
		if env.HasRessourcesToSellAtHarbour(ship) {
			ship.SetAction("sell")
			return
		} else {
			//Find harbour that is good for money ching ching
			//pathfinding

			ship.SetAction("follow")
			return
		}

	}

	ship.SetAction("follow")
	return

}

func merchantShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Merchant ship act...")
	// Merchant ship can trade with port to level up but can't attack or defend itself
}

func merchantShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("Merchant ship level up...")
	// merchantShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Privateer ship -------------------

func privateerShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("Privateer ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func privateerShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("Privateer ship deliberate...")
}

func privateerShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Privateer ship act...")
	// Privateer ship can attack Pirate ship, Merchant ship, Filibusteer ship, Army ship
}

func privateerShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("Privateer ship level up...")
	// privateerShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Army ship -------------------

func armyShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("Army ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func armyShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("Army ship deliberate...")
}

func armyShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Army ship act...")
	// Army ship can attack Pirate ship,  Filibusteer ship, Privateer ship
}

func armyShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("Army ship level up is not possible...")
	// Army ship can't level up i think

}

// ------------------- Filibusteer ship -------------------

func filibusteerShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("Filibusteer ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func filibusteerShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("Filibusteer ship deliberate...")
}

func filibusteerShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Filibusteer ship act...")
}

func filibusteerShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("Filibusteer ship level up...")
	// filibusteerpShipLevelUp(ship *Ship, env *Environment)
}

// ------------------- Pirate hunter ship -------------------

func piratehunterShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("PirateHunter ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func piratehunterShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("PirateHunter ship deliberate...")
}

func piratehunterShipAct(ship *Ship, env *Environment) {
	//fmt.Println("PirateHunter ship act...")
	// PirateHunter ship can attack Pirate ship, Filibusteer ship, Privateer ship

}

func piratehunterShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("PirateHunter ship level up...")
	// piratehunterShipLevelUp(ship *Ship, env *Environment)
}

// ------------------- Zoro ship -------------------

func zoroShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("Zoro ship percept...")
	// //fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby
}

func zoroShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("Zoro ship deliberate...")
}

func zoroShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Zoro ship act...")
	//LostZoro(ship, env)

}

func zoroShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("Zoro ship level up is not possible...")
	// no level up for zoro ship, can't access to the port
}

// ------------------- Rackham ship -------------------

func rackhamShipPercept(ship *Ship, env *Environment) {
	//fmt.Println("Rackham ship percept...")
}

func rackhamShipDeliberate(ship *Ship, env *Environment) {
	//fmt.Println("Rackham ship deliberate...")
}

func rackhamShipAct(ship *Ship, env *Environment) {
	//fmt.Println("Rackham ship act...")

}

func rackhamShipLevelUp(ship *Ship, env *Environment) {
	//fmt.Println("Rackham ship level up is not possible...")
}
