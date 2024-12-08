package simulation

import (
	"fmt"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
)

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

func armyShipDeliberate(ship *Ship, env *Environment) {
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

func armyShipAct(ship *Ship, env *Environment) {
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
