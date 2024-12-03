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
	market    Market
	agentLock sync.RWMutex
}

// NewEnvironment initializes the Environment
func NewEnvironment(agents []Agent, objects []Object) (env *Environment) {
	market := NewMarket()
	return &Environment{
		agents:  agents,
		objects: objects,
		market:  market,
	}
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

// Remove the Agent from the Environment
func (env *Environment) RemoveAgent(target Agent) {
	// Loop over the agents until the target is found
	for i, agent := range env.agents {
		if agent == target {
			// Here we take all the agents except the one to remove it
			env.agents = append(env.agents[:i], env.agents[i+1:]...)
			return
		}
	}
}

/*
	Perception of the environment by the agents
*/

// Basic Euclidian Distance [[Need to change the distance use so the world is round and there is terrain where agents can't go]]
func CalculateDistance(p1, p2 Coordinates) float64 {
	return math.Sqrt(math.Pow(float64(p2.X())-float64(p1.X()), 2) + math.Pow(float64(p2.Y())-float64(p1.Y()), 2))
}

// AgentsNearby returns all agents in the fov of the current agent
func (env *Environment) AgentsNearby(agt Agent) []Agent {
	// Initialize the nearby agents array
	nearbyAgents := []Agent{}

	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()

	// Loop on all agents in the environment
	for _, otherAgent := range env.agents {

		// Skip the current agent
		if otherAgent == agt {
			continue
		}

		// Calculate the distance between the agents [[CHANGE THIS WITH THE FORMULA WE USE]]
		otherAgtCoord := otherAgent.Coordinates()
		distance := CalculateDistance(otherAgtCoord, agtCoord)

		// Check if the other agent is within the field of view
		if int(distance) <= agtFov {
			nearbyAgents = append(nearbyAgents, otherAgent)
		}
	}

	return nearbyAgents
}

// KrakenNearby returns true if the kraken is within the field of view of the ship
func (env *Environment) KrakenNearby(agt Agent) Agent {
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
		if otherAgent.AgtType() == "kraken" {
			// Calculate the distance between the agents
			otherAgtCoord := otherAgent.Coordinates()
			distance := CalculateDistance(otherAgtCoord, agtCoord)

			// Check if the other agent is within the field of view
			if int(distance) <= agtFov {
				return otherAgent
			}
		}

	}

	return nil
}

// ArmyNearby returns true if the army is within the field of view of the ship
func (env *Environment) ArmyNearby(agt Agent) Agent {
	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()

	// Loop on all agents in the environment
	for _, otherAgent := range env.agents {

		// Skip the current agent
		if otherAgent != agt {
			continue // Kraken already used
		}

		if otherAgent.AgtType() == "ship" {
			// check if the ship is docked and if it is an army ship
			if ship, ok := otherAgent.(*Ship); ok && !ship.Docked() {
				if ship.ShipType() == "army" {
					// Calculate the distance between the agents
					otherAgtCoord := otherAgent.Coordinates()
					distance := CalculateDistance(otherAgtCoord, agtCoord)

					// Check if the other agent is within the field of view
					if int(distance) <= agtFov {
						return otherAgent
					}

				}
			}
		}
		return nil
	}
}

// PirateNearby returns true if the pirate is within the field of view of the ship
func (env *Environment) PirateNearby(agt Agent) Agent {
	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()

	// Loop on all agents in the environment
	for _, otherAgent := range env.agents {

		// Skip the current agent
		if otherAgent != agt {
			continue // Kraken already used
		}

		if otherAgent.AgtType() == "ship" {
			// check if the ship is docked and if it is an pirate ship
			if ship, ok := otherAgent.(*Ship); ok && !ship.Docked() {
				if ship.ShipType() == "pirate" {
					// Calculate the distance between the agents
					otherAgtCoord := otherAgent.Coordinates()
					distance := CalculateDistance(otherAgtCoord, agtCoord)

					// Check if the other agent is within the field of view
					if int(distance) <= agtFov {
						return otherAgent
					}

				}
			}
		}
		return nil
	}
}

// ZoroNearby returns true if the Zoro is within the field of view of the ship
func (env *Environment) ZoroNearby(agt Agent) Agent {
	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()

	// Loop on all agents in the environment
	for _, otherAgent := range env.agents {

		// Skip the current agent
		if otherAgent != agt {
			continue // Kraken already used
		}

		if otherAgent.AgtType() == "ship" {
			// check if the ship is docked and if it is an zoro ship
			if ship, ok := otherAgent.(*Ship); ok && !ship.Docked() {
				if ship.ShipType() == "zoro" {
					// Calculate the distance between the agents
					otherAgtCoord := otherAgent.Coordinates()
					distance := CalculateDistance(otherAgtCoord, agtCoord)

					// Check if the other agent is within the field of view
					if int(distance) <= agtFov {
						return otherAgent
					}

				}
			}
		}
		return nil
	}
}

// MerchantNearby returns true if the army is within the field of view of the ship
func (env *Environment) MerchantNearby(agt Agent) Agent {
	// Get the coordinates and FOV of the agent
	agtCoord := agt.Coordinates()
	agtFov := agt.Fov()

	// Loop on all agents in the environment
	for _, otherAgent := range env.agents {

		// Skip the current agent
		if otherAgent != agt {
			continue // Kraken already used
		}

		if otherAgent.AgtType() == "ship" {
			// check if the ship is docked and if it is an merchant ship
			if ship, ok := otherAgent.(*Ship); ok && !ship.Docked() {
				if ship.ShipType() == "merchant" {
					// Calculate the distance between the agents
					otherAgtCoord := otherAgent.Coordinates()
					distance := CalculateDistance(otherAgtCoord, agtCoord)

					// Check if the other agent is within the field of view
					if int(distance) <= agtFov {
						return otherAgent
					}

				}
			}
		}
		return nil
	}
}

// ClosestAgent returns the closest agent to the given agent that is not docked within it's fov
func (env *Environment) ClosestAgent(agt Agent) Agent {
	var closestAgent Agent
	minDistance := math.MaxFloat64

	// Get the coordinates of the agent
	agtCoord := agt.Coordinates()

	// Loop through all agents in the environment
	for _, otherAgent := range env.agents {
		// Skip the current agent
		if otherAgent.Id() == agt.Id() {
			continue
		}

		// Skip docked ships because they are safe
		if otherAgent.AgtType() == "ship" {
			if ship, ok := otherAgent.(*Ship); ok && ship.Docked() {
				continue
			}
		}

		// Calculate the distance between the agents
		otherAgtCoord := otherAgent.Coordinates()
		distance := CalculateDistance(otherAgtCoord, agtCoord)

		// If the Agent if further than the Fov of the Agent, then we don't return it
		if distance > float64(agt.Fov()) {
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
