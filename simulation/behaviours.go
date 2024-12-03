package simulation

import "fmt"

func pirateShipPercept(ship *Ship, env *Environment) {
	// If kraken nearby set danger to kraken
	kraken := env.KrakenNearby(ship)
	if kraken != nil {
		ship.SetDanger(kraken)
	} else {
		ship.SetDanger(nil)

	}

	// If army nearby set danger to pirate
	army := env.ArmyNearby(ship)
	if army != nil {
		ship.SetDanger(army) //problème car je reçois une liste d'agents et pas un seul
	} else {
		ship.SetDanger(nil)
	}

}

func pirateShipDeliberate(ship *Ship) {
	fmt.Println("Pirate ship deliberate...")

}

func pirateShipAct(ship *Ship, env *Environment) {
	fmt.Println("Pirate ship act...")
	// Pirate ship can attack Merchant ship, Filibusteer ship, Privateer ship, Army ship
}

func pirateShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("Pirate ship level up...")
	// pirateShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Merchant ship -------------------

func merchantShipPercept(ship *Ship, env *Environment) {
	fmt.Println("Merchant ship percept...")
	// fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func merchantShipDeliberate(ship *Ship) {
	fmt.Println("Merchant ship deliberate...")
}

func merchantShipAct(ship *Ship, env *Environment) {
	fmt.Println("Merchant ship act...")
	// Merchant ship can trade with port to level up but can't attack or defend itself
}

func merchantShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("Merchant ship level up...")
	// merchantShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Privateer ship -------------------

func privateerShipPercept(ship *Ship, env *Environment) {
	fmt.Println("Privateer ship percept...")
	// fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func privateerShipDeliberate(ship *Ship) {
	fmt.Println("Privateer ship deliberate...")
}

func privateerShipAct(ship *Ship, env *Environment) {
	fmt.Println("Privateer ship act...")
	// Privateer ship can attack Pirate ship, Merchant ship, Filibusteer ship, Army ship
}

func privateerShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("Privateer ship level up...")
	// privateerShipLevelUp(ship *Ship, env *Environment) //regarde si elle est good pour level up
}

// ------------------- Army ship -------------------

func armyShipPercept(ship *Ship, env *Environment) {
	fmt.Println("Army ship percept...")
	// fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func armyShipDeliberate(ship *Ship) {
	fmt.Println("Army ship deliberate...")
}

func armyShipAct(ship *Ship, env *Environment) {
	fmt.Println("Army ship act...")
	// Army ship can attack Pirate ship,  Filibusteer ship, Privateer ship
}

func armyShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("Army ship level up is not possible...")
	// Army ship can't level up i think

}

// ------------------- Filibusteer ship -------------------

func filibusteerShipPercept(ship *Ship, env *Environment) {
	fmt.Println("Filibusteer ship percept...")
	// fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func filibusteerShipDeliberate(ship *Ship) {
	fmt.Println("Filibusteer ship deliberate...")
}

func filibusteerShipAct(ship *Ship, env *Environment) {
	fmt.Println("Filibusteer ship act...")
}

func filibusteerShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("Filibusteer ship level up...")
	// filibusteerpShipLevelUp(ship *Ship, env *Environment)
}

// ------------------- Pirate hunter ship -------------------

func piratehunterShipPercept(ship *Ship, env *Environment) {
	fmt.Println("PirateHunter ship percept...")
	// fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby

}

func piratehunterShipDeliberate(ship *Ship) {
	fmt.Println("PirateHunter ship deliberate...")
}

func piratehunterShipAct(ship *Ship, env *Environment) {
	fmt.Println("PirateHunter ship act...")
	// PirateHunter ship can attack Pirate ship, Filibusteer ship, Privateer ship

}

func piratehunterShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("PirateHunter ship level up...")
	// piratehunterShipLevelUp(ship *Ship, env *Environment)
}

// ------------------- Zoro ship -------------------

func zoroShipPercept(ship *Ship, env *Environment) {
	fmt.Println("Zoro ship percept...")
	// fmt.Println(AgentsNearby(ag.ship)) // AgentsNearby is a function that returns the agents nearby
}

func zoroShipDeliberate(ship *Ship) {
	fmt.Println("Zoro ship deliberate...")
}

func zoroShipAct(ship *Ship, env *Environment) {
	fmt.Println("Zoro ship act...")
	//LostZoro(ship, env)

}

func zoroShipLevelUp(ship *Ship, env *Environment) {
	fmt.Println("Zoro ship level up is not possible...")
	// no level up for zoro ship, can't access to the port
}
