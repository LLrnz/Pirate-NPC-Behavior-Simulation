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
	sync.RWMutex
	agents    []Agent
	objects   []Object
	harbours  []Harbour
	market    *Market
	agentLock sync.RWMutex
}

// NewEnvironment initializes the Environment
func NewEnvironment(agents []Agent, objects []Object, harbours []Harbour) (env *Environment) {
	market := NewMarket()
	return &Environment{
		agents:   agents,
		objects:  objects,
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
	return env.agents
}

// Add an Agent to the Environment
func (env *Environment) AddAgent(agt Agent) {
	env.agents = append(env.agents, agt)
}

// Add an Object to the Environment
func (env *Environment) AddObject(object Object) {
	env.objects = append(env.objects, object)
}

// Add an Harbour to the Environment
func (env *Environment) AddHarbour(h Harbour) {
	env.harbours = append(env.harbours, h)
}

// Remove the Agent from the Environment
func (env *Environment) RemoveAgent(target Agent) {
	// Loop over the agents until the target is found
	for i, agent := range env.agents {
		if agent.Id() == target.Id() {
			// Here we take all the agents except the one to remove it
			env.agents = append(env.agents[:i], env.agents[i+1:]...)
			return
		}
	}
}

// Attack the agent in the env
func (env *Environment) AttackAgent(target Agent, attacker Agent) {
	// Loop over the agents until the target is found
	for _, agent := range env.agents {
		if agent.Id() == target.Id() {
			agent.RemoveHealth(attacker.Attack())

			if agent.Health() <= 0 {
				env.RemoveAgent(target)
			}
			// If die, retrieve 60% ressources
			return
		}
	}
}

/*
	Perception of the environment by the agents
*/

// KrakenNearby returns true if the kraken is within the field of view of the agent
func (env *Environment) KrakenNearby(agt Agent) *Agent {
	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()

	// Loop on all agents in the environment
	for _, otherAgent := range env.agents {

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
func (env *Environment) ClosestShipFromKraken(agt *Kraken) *Agent {
	var closestAgent *Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := agt.Coordinates()

	// Loop through all agents in the environment
	for _, otherAgent := range env.agents {
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
			closestAgent = &otherAgent
		}
	}

	return closestAgent
}

// Get harbours in the environment
func (env *Environment) Harbours() []Harbour {
	return env.harbours
}

// IsHarbour checks if the ship's current position matches any harbor's coordinates
func (env *Environment) IsHarbour(ship *Ship) (*Harbour, bool) {
	// Get the ship coordinates
	shipCoordinates := ship.Coordinates()

	// Loop through harbours
	for _, harbour := range env.Harbours() {
		if harbour.Coordinates() == shipCoordinates {
			return &harbour, true
		}
	}

	// Not a harbour
	return nil, false
}

// DockShip docks a ship at the specified harbor
func (env *Environment) DockShip(ship *Ship, harbour *Harbour) {
	// Set the ship as docked at the harbor
	ship.SetDocked(harbour)

	// Add the ship to docked agents
	harbour.DockShip(ship)
}

// UndockShip undocks a ship at it's harbour
func (env *Environment) UndockShip(ship *Ship) {
	// Remove the ship to docked agents
	ship.Docked().UndockShip(ship)

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
				print(ressource)
				return true
			}
		}
	}

	return false
}

// ClosestPirate returns the closest pirate ship not docked within it's fov
func (env *Environment) ClosestPirate(ag *Ship) *Agent {
	var closestAgent *Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()

	// Loop through all agents in the environment
	for _, otherAgent := range env.agents {
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
			closestAgent = &otherAgent
		}
	}

	return closestAgent
}

// ClosestPrivateer returns the closest privateer ship not docked within it's fov and not same nationality
func (env *Environment) ClosestPrivateerNationality(ag *Ship) *Agent {
	var closestAgent *Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := ag.Coordinates()

	// Loop through all agents in the environment
	for _, otherAgent := range env.agents {
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
			closestAgent = &otherAgent
		}
	}

	return closestAgent
}

// ClosestShip returns the closest ship to the given ship that is not docked within it's fov
func (env *Environment) ClosestShipFromShip(agt *Ship) *Agent {
	var closestAgent *Agent
	minDistance := math.MaxInt

	// Get the coordinates of the agent
	agtCoord := agt.Coordinates()

	// Loop through all agents in the environment
	for _, otherAgent := range env.agents {
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
			closestAgent = &otherAgent
		}
	}

	return closestAgent
}
