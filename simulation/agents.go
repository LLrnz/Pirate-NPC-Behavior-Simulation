package simulation

import (
	"log"
	"math"
)

/*
********************************************
		agents.go - Agent logique
********************************************
*/

/*
----------------------------------------
*********** Agent Interface ************
----------------------------------------
*/
// An Agent is an Object

type Agent interface {
	Object

	Percept(env *Environment)
	Deliberate()
	Act(env *Environment)

	Fov() int
	Id() string
	AgtType() string
}

/*
---------------------------------------
************* BaseAgent ***************
---------------------------------------
*/
// It is the agent struct that is used by any agent. It uses a BaseObject struct

type BaseAgent struct {
	BaseObject
	id      string
	fov     int
	agtType string
}

// Getter Field of view
func (agt *BaseAgent) Fov() int {
	return agt.fov
}

// Getter Id
func (agt *BaseAgent) Id() string {
	return agt.id
}

// Getter Agent type
func (agt *BaseAgent) AgtType() string {
	return agt.agtType
}

func (agt *BaseAgent) Percept(env *Environment) {
	// Base implementation
}

func (agt *BaseAgent) Deliberate() {
	// Base implementation
}

func (agt *BaseAgent) Act(env *Environment) {
	// Base implementation
}

/*
--------------------------------------
*************** Ship *****************
--------------------------------------
*/
// There are different types of ships [[Explain here what behaviours they will have]]

type Ship struct {
	*BaseAgent
	speed               int
	ressources          map[string]int
	money               float32
	docked              bool
	shipType            string
	capacity            int
	attack              int
	danger              Agent
	behaviourPercept    func(*Ship, *Environment)
	behaviourDeliberate func(*Ship)
	behaviourAct        func(*Ship, *Environment)
	behaviourLevelUp    func(*Ship, *Environment)
}

// Different types of ships
// Pirate, Merchant, Privateer, Army, Filibusteer, PirateHunter, Zoro

// Pirate ship
func NewShipPirate(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "pirate",
		capacity:            10,
		attack:              20,
		behaviourPercept:    pirateShipPercept,
		behaviourDeliberate: pirateShipDeliberate,
		behaviourAct:        pirateShipAct,
		behaviourLevelUp:    pirateShipLevelUp,
	}
}

// Merchant ship
func NewShipMerchant(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "merchant",
		capacity:            100,
		attack:              2,
		behaviourPercept:    merchantShipPercept,
		behaviourDeliberate: merchantShipDeliberate,
		behaviourAct:        merchantShipAct,
		behaviourLevelUp:    merchantShipLevelUp,
	}
}

// Privateer ship
func NewShipPrivateer(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "privateer",
		capacity:            10,
		attack:              20,
		behaviourPercept:    privateerShipPercept,
		behaviourDeliberate: privateerShipDeliberate,
		behaviourAct:        privateerShipAct,
		behaviourLevelUp:    privateerShipLevelUp,
	}
}

// Army ship
func NewShipArmy(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "army",
		behaviourPercept:    armyShipPercept,
		behaviourDeliberate: armyShipDeliberate,
		behaviourAct:        armyShipAct,
		behaviourLevelUp:    armyShipLevelUp,
	}
}

// Filibusteer ship
func NewShipFilibusteer(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "filibusteer",
		capacity:            30,
		attack:              6,
		behaviourPercept:    filibusteerShipPercept,
		behaviourDeliberate: filibusteerShipDeliberate,
		behaviourAct:        filibusteerShipAct,
		behaviourLevelUp:    filibusteerShipLevelUp,
	}
}

// Pirate hunter ship
func NewShipPirateHunter(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "piratehunter",
		capacity:            10,
		attack:              30,
		behaviourPercept:    piratehunterShipPercept,
		behaviourDeliberate: piratehunterShipDeliberate,
		behaviourAct:        piratehunterShipAct,
		behaviourLevelUp:    piratehunterShipLevelUp,
	}
}

// Zoro ship
func NewShipZoro(id string, fov int, cor Coordinates, docked bool, env *Environment) *Ship {
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "ship",
		},
		docked:              docked,
		speed:               1,
		shipType:            "zoro",
		capacity:            10,
		attack:              200,
		behaviourPercept:    zoroShipPercept,
		behaviourDeliberate: zoroShipDeliberate,
		behaviourAct:        zoroShipAct,
		behaviourLevelUp:    zoroShipLevelUp,
	}
}

func (ag *Ship) MoveTowards(targetCoords Coordinates) {
	current := ag.Coordinates()

	// Determine step size based on Ship speed
	dx := targetCoords.X() - current.X()
	dy := targetCoords.Y() - current.Y()
	stepX, stepY := 0, 0

	if dx != 0 {
		stepX = ag.speed * dx / int(math.Abs(float64(dx)))
	}
	if dy != 0 {
		stepY = ag.speed * dy / int(math.Abs(float64(dy)))
	}

	// Update Kraken coordinates
	ag.SetCoordinates(Coordinates{
		x: current.X() + stepX,
		y: current.Y() + stepY,
	})
}

func (ag *Ship) Percept(env *Environment) {
	ag.behaviourPercept(ag, env)
}

func (ag *Ship) Deliberate() {
	ag.behaviourDeliberate(ag)
}

func (ag *Ship) Act(env *Environment) {
	ag.behaviourAct(ag, env)
}

func (ag *Ship) LevelUp(env *Environment) {
	ag.behaviourLevelUp(ag, env)
}

func (ag *Ship) Docked() bool {
	return ag.docked
}

func (ag *Ship) ShipType() string {
	return ag.shipType
}

func (ag *Ship) SetDanger(danger Agent) {
	ag.danger = danger
}

func (ag *Ship) Danger() Agent {

/*
--------------------------------------
*********** Kraken Agent *************
--------------------------------------
*/
// The Kraken finds closest agents and attacks them

type Kraken struct {
	*BaseAgent
	health int
	attack int
	speed  int
	target Agent
	action string
}

func NewKraken(id string, fov int, attack int, speed int, cor Coordinates) *Kraken {
	return &Kraken{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: cor},
			id:         id,
			fov:        fov,
			agtType:    "kraken",
		},
		attack: attack,
		speed:  speed,
		health: 100,
	}
}

func (ag *Kraken) RestAndGainHealth() {
	ag.health = ag.health + 10
}

func (ag *Kraken) MoveTowards(targetCoords Coordinates) {
	current := ag.Coordinates()

	// Determine step size based on Kraken's speed
	dx := targetCoords.X() - current.X()
	dy := targetCoords.Y() - current.Y()
	stepX, stepY := 0, 0

	if dx != 0 {
		stepX = ag.speed * dx / int(math.Abs(float64(dx)))
	}
	if dy != 0 {
		stepY = ag.speed * dy / int(math.Abs(float64(dy)))
	}

	// Update Kraken coordinates
	ag.SetCoordinates(Coordinates{
		x: current.X() + stepX,
		y: current.Y() + stepY,
	})
}

func (ag *Kraken) Percept(env *Environment) {
	// Retrieve the closest agent from kraken
	closest := env.ClosestAgent(ag)

	// Check if there is a closest agent
	if closest != nil {
		ag.target = closest
	} else {
		ag.target = nil
	}
}

func (ag *Kraken) Deliberate() {

	//If the Kraken is low life or has no target = rest to gain PV
	if ag.health <= 20 || ag.target == nil {
		ag.action = "rest"
		return
	}

	// Calculate the distance to the target agent
	distance := CalculateDistance(ag.Coordinates(), ag.target.Coordinates())

	// If the Kraken can attack, it should attack the ship
	// Else it should follow the nearest agent as a target
	if distance <= float64(ag.speed) {
		ag.action = "attack"
	} else {
		ag.action = "follow"
	}
}

func (ag *Kraken) Act(env *Environment) {
	// No action = do nothing
	if ag.action == "" {
		return
	}

	switch ag.action {
	case "attack":
		// Remove the target from the environment (we should only remove health points)
		env.RemoveAgent(ag.target)
	case "follow":
		// Move towards target
		ag.MoveTowards(ag.target.Coordinates())
	case "rest":
		// Regain 10% of health points
		ag.RestAndGainHealth()
	default:
		log.Printf("%s has no valid action to perform.\n", ag.id)
	}

	ag.target = nil
	ag.action = ""
}
