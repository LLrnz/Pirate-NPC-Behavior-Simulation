package simulation

import (
	"math"
	"sync"
)

/*
--------------------------------------
*********** Environment **************
--------------------------------------
*/
// The Environment handles all integractions between objects

type Environment struct {
	agents   []Agent
	harbours []*Harbour
	market   *Market

	sync.RWMutex
}

// NewEnvironment initializes the Environment
func NewEnvironment(agents []Agent, harbours []*Harbour) (env *Environment) {
	market := NewMarket()
	return &Environment{
		agents:   agents,
		harbours: harbours,
		market:   market,
	}
}

// Getter that returns the Market
func (env *Environment) Market() *Market {
	return env.market
}

// Getter that returns the Agents
func (env *Environment) Agents() []Agent {
	env.RLock()
	defer env.RUnlock()

	return env.agents
}

// Add an Agent to the Environment
func (env *Environment) AddAgent(agt Agent) {
	env.Lock()
	defer env.Unlock()

	env.agents = append(env.agents, agt)
}

// Remove the Agent from the Environment
func (env *Environment) RemoveAgent(target Agent) {
	agents := env.Agents()

	env.Lock()
	defer env.Unlock()

	// Loop over the agents until the target is found
	for i, agent := range agents {
		if agent.Id() == target.Id() {
			// Here we take all the agents except the one to remove it
			env.agents = append(env.agents[:i], env.agents[i+1:]...)
			return
		}
	}
}

// Get harbours in the environment
func (env *Environment) Harbours() []*Harbour {
	env.RLock()
	defer env.RUnlock()

	return env.harbours
}

// Add an Harbour to the Environment
func (env *Environment) AddHarbour(h *Harbour) {
	env.Lock()
	defer env.Unlock()

	env.harbours = append(env.harbours, h)
}

/*
	Perception of the environment by the agents
*/

// IsHarbour checks if the ship's current position matches any harbor's coordinates
func (env *Environment) IsHarbour(ship *Ship) (*Harbour, bool) {
	// Get the ship coordinates
	shipCoordinates := ship.Coordinates()
	harbours := env.Harbours()

	env.RLock()
	defer env.RUnlock()

	// Loop through harbours
	for _, harbour := range harbours {
		if harbour.Coordinates() == shipCoordinates {
			return harbour, true
		}
	}

	// Not a harbour
	return nil, false
}

// ClosestHarbour returns the closest harbour
func (env *Environment) ClosestHarbour(ag *Ship) *Harbour {
	var closestHarbour *Harbour
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	harbours := env.Harbours()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, harbour := range harbours {
		// Calculate the distance between the kraken and the ship
		otherAgtCoord := harbour.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestHarbour = harbour
		}
	}

	return closestHarbour
}

// DockShip docks a ship at the specified harbor
func (env *Environment) DockShip(ship *Ship, harbour *Harbour) {
	// Set the ship as docked at the harbor
	ship.SetDocked(harbour)
}

// UndockShip undocks a ship at it's harbour
func (env *Environment) UndockShip(ship *Ship) {
	// Remove docked ship status
	ship.SetDocked(nil)
}

// HasRessourcesToSellAtHarbour returns true if ship has ressources to sell else false
func (env *Environment) HasRessourcesToSellAtHarbour(agt *Ship) bool {
	// Get Harbour where the ship is docked
	harbour := agt.Docked()

	// Get the ressources the ship has
	agtRessources := agt.Ressources()

	// For each ressource of the ship, check if it can be sold
	for ressource, quantity := range agtRessources {
		if quantity != 0 {
			if harbour.IsRessourcesSellable(ressource) {
				return true
			}
		}
	}

	return false
}

// KrakenNearby returns true if the kraken is within the field of view of the agent
func (env *Environment) KrakenNearby(agt Agent) *Agent {
	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop on all agents in the environment
	for _, otherAgent := range agents {

		// Skip the current agent
		if otherAgent == agt {
			continue
		}

		// If kraken agent
		if otherAgent.AgtType() == "Kraken" {
			// Calculate the distance between the agents
			otherAgtCoord := otherAgent.Coordinates()
			distance := Heuristic(otherAgtCoord, agtCoord)

			// Check if the other agent is within the field of view
			if int(distance) <= agtFov {
				return &otherAgent
			}
		}

	}

	return nil
}

// ClosestShip returns the closest ship to the given ship that is not docked within it's fov
func (env *Environment) ClosestShipFromKraken(agt *Kraken) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := agt.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == agt.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		//print("Ship ", ship.Id(), "docked ", ship.Docked())
		if ok && ship.Docked() != nil {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > agt.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestPirate returns the closest pirate ship not docked within it's fov
func (env *Environment) ClosestPirate(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Pirate") {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestPrivateer returns the closest privateer ship not docked within it's fov and not same nationality
func (env *Environment) ClosestPrivateerNationality(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Privateer" || ship.Nationality() == ag.Nationality()) {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestShip returns the closest ship to the given ship that is not docked within it's fov
func (env *Environment) ClosestShipFromShip(agt *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := agt.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == agt.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && ship.Docked() != nil {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > agt.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestMerchant returns the closest merchant ship not docked within it's fov
func (env *Environment) ClosestMerchant(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Merchant") {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestMerchantNationality returns the closest army ship not docked within it's fov and not same nationality
func (env *Environment) ClosestMerchantNationality(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Merchant" || ship.Nationality() == ag.Nationality()) {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestPrivateer returns the closest privateer ship not docked within it's fov
func (env *Environment) ClosestPrivateer(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Privateer") {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestPrivateer returns the closest army ship not docked within it's fov
func (env *Environment) ClosestArmy(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Army") {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// ClosestArmyNationality returns the closest army ship not docked within it's fov and not same nationality
func (env *Environment) ClosestArmyNationality(ag *Ship) Agent {
	var closestAgent Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()
	agents := env.Agents()

	env.RLock()
	defer env.RUnlock()

	// Loop through all agents in the environment
	for _, otherAgent := range agents {
		// Skip the current agent
		if otherAgent.Id() == ag.Id() || otherAgent.AgtType() != "Ship" {
			continue
		}

		// Skip docked ships because they are safe
		ship, ok := otherAgent.(*Ship)
		if ok && (ship.Docked() != nil || ship.ShipType() != "Army" || ship.Nationality() == ag.Nationality()) {
			continue
		}

		// Calculate the distance between the kraken and the ship
		otherAgtCoord := otherAgent.Coordinates()
		distance := Heuristic(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > ag.Fov() {
			continue
		}

		// Update the closest agent if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestAgent = otherAgent
		}
	}

	return closestAgent
}

// Attack the agent in the env
func (env *Environment) AttackAgent(target Agent, attacker Agent) {
	agents := env.Agents()

	env.Lock()
	defer env.Unlock()

	// Loop over the agents until the target is found
	for i, agent := range agents {
		if agent.Id() == target.Id() {
			agent.RemoveHealth(attacker.Attack())

			if agent.Health() <= 0 {
				ship, ok := agent.(*Ship)
				if ok && ship.Quantity() != 0 {
					// Get target ressources
					currentRessources := ship.Ressources()
					modifiedRessources := make(map[string]int)

					// Only get 60% of what it had because it burnt
					for ressource, quantity := range currentRessources {
						modifiedQuantity := int(float64(quantity) * 0.6)
						modifiedRessources[ressource] = modifiedQuantity
					}

					shipAttacker, ok := attacker.(*Ship)
					if ok {

						shipAttacker.SetRessources(modifiedRessources)
					}
				}

				env.agents = append(env.agents[:i], env.agents[i+1:]...)
			}
			return
		}
	}
}
