package simulation

import (
	"sync"
	"time"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
	"golang.org/x/exp/rand"
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
	Id() string
	AgtType() string

	Fov() int
	SetFov(fov int)
	Speed() int
	SetSpeed(speed int)
	Attack() int
	SetAttack(attack int)
	Health() int
	SetHealth(health int)
	AddHealth(health int)
	RemoveHealth(health int)

	Target() Agent
	SetTarget(target Agent)
	Action() string
	SetAction(action string)
	Itinerary() []Coordinates
	SetItinerary(itinerary []Coordinates)
	ResetItinerary()

	FollowItinerary(mistake bool)

	Percept(env *Environment, step int)
	Deliberate(env *Environment, step int)
	Act(env *Environment, step int)
}

/*
---------------------------------------
************* BaseAgent ***************
---------------------------------------
*/
// It is the agent struct that is used by any agent. It uses a BaseObject struct

type BaseAgent struct {
	BaseObject
	sync.RWMutex
	id      string
	agtType string

	fov    int
	speed  int
	attack int
	health int

	target    Agent
	action    string
	itinerary []Coordinates
	direction Direction
}

// Getter Id
func (agt *BaseAgent) Id() string {
	return agt.id
}

// Getter Agent type (Ship or Kraken)
func (agt *BaseAgent) AgtType() string {
	return agt.agtType
}

// Getter Field of view
func (agt *BaseAgent) Fov() int {
	agt.RLock()
	defer agt.RUnlock()

	return agt.fov
}

// Setter Fov
func (agt *BaseAgent) SetFov(fov int) {
	agt.fov = fov
}

// Getter Speed
func (agt *BaseAgent) Speed() int {
	return agt.speed
}

// Setter Speed
func (agt *BaseAgent) SetSpeed(speed int) {
	agt.speed = speed
}

// Getter Attack
func (agt *BaseAgent) Attack() int {
	agt.RLock()
	defer agt.RUnlock()

	return agt.attack
}

// Setter Attack
func (agt *BaseAgent) SetAttack(attack int) {
	agt.Lock()
	defer agt.Unlock()

	agt.attack = attack
}

// Getter Health
func (agt *BaseAgent) Health() int {
	agt.RLock()
	defer agt.RUnlock()

	return agt.health
}

// Setter Health
func (agt *BaseAgent) SetHealth(health int) {
	agt.Lock()
	defer agt.Unlock()

	agt.health = health
}

// Setter Health Add
func (agt *BaseAgent) AddHealth(health int) {
	agt.Lock()
	defer agt.Unlock()

	agt.health = agt.health + health
}

// Setter Health Remove
func (agt *BaseAgent) RemoveHealth(health int) {
	agt.Lock()
	defer agt.Unlock()

	agt.health = agt.health - health
}

// Getter Target
func (agt *BaseAgent) Target() Agent {
	return agt.target
}

// Setter Target
func (agt *BaseAgent) SetTarget(target Agent) {
	agt.target = target
}

// Getter Action
func (agt *BaseAgent) Action() string {
	return agt.action
}

// Setter Action
func (agt *BaseAgent) SetAction(action string) {
	agt.action = action
}

// Getter Itinerary
func (agt *BaseAgent) Itinerary() []Coordinates {
	return agt.itinerary
}

// Setter Itinerary
func (agt *BaseAgent) SetItinerary(itinerary []Coordinates) {
	agt.itinerary = itinerary
}

func (ag *BaseAgent) ResetItinerary() {
	ag.SetItinerary([]Coordinates{})
}

// Getter Direction
func (agt *BaseAgent) Direction() Direction {
	return agt.direction
}

// Setter Direction
func (agt *BaseAgent) SetDirection(direction Direction) {
	agt.direction = direction
}

// Compute current Direction
func (agt *BaseAgent) CurrentDirection() {
	if len(agt.Itinerary()) == 0 || agt.Itinerary() == nil {
		return
	}

	if agt.Itinerary()[0].X() > agt.Coordinates().x {
		agt.SetDirection(Right)
		return
	}

	if agt.Itinerary()[0].X() < agt.Coordinates().x {
		agt.SetDirection(Left)
		return
	}

	if agt.Itinerary()[0].Y() > agt.Coordinates().y {
		agt.SetDirection(Up)
		return
	}

	if agt.Itinerary()[0].Y() < agt.Coordinates().y {
		agt.SetDirection(Down)
		return
	}
}

// Agent Percept
func (agt *BaseAgent) Percept(env *Environment) {

}

// Agent Deliberate
func (agt *BaseAgent) Deliberate() {

}

// Agent Act
func (agt *BaseAgent) Act(env *Environment) {

}

// FollowItinerary at speed of agent
func (ag *BaseAgent) FollowItinerary(mistake bool) {

	if mistake == false {

		itinerary := ag.Itinerary()
		speed := ag.Speed()

		// If there is not enough coordinates to go through
		if len(itinerary) == 0 {
			return
		}

		// Move based on the Agent speed
		moves := speed
		if moves > len(itinerary) {
			moves = len(itinerary) // Max to the end of the array
		}

		// Update the Agent coordinates to the last point within the move range
		newCoordinates := itinerary[moves-1]
		ag.SetCoordinates(newCoordinates)

		// Remove the used coordinates from the itinerary
		ag.SetItinerary(itinerary[moves:])
	} else {
		// Get current coordinates
		currentCoordinates := ag.Coordinates()

		// Get neighbors
		neighbors := GetNeighbors(currentCoordinates, constants.GridWidth, constants.GridHeight)

		// Filter neighbors that are navigable
		var passableNeighbors []Coordinates
		for _, neighbor := range neighbors {
			if constants.MapGrid[neighbor.Y()][neighbor.X()] != 1 {
				passableNeighbors = append(passableNeighbors, neighbor)
			}
		}

		// Select a random valid neighbor
		if len(passableNeighbors) > 0 {
			rand.Seed(uint64(time.Now().UnixNano()))
			randomIndex := rand.Intn(len(passableNeighbors))
			randomNeighbor := passableNeighbors[randomIndex]

			ag.SetCoordinates(randomNeighbor)
		}
	}
}

/*
--------------------------------------
*************** Ship *****************
--------------------------------------
*/
// There are different types of ships [[Explain here what behaviours they will have]]

type Ship struct {
	*BaseAgent

	sync.RWMutex
	shipType    string
	nationality constants.Nationality
	shipState   ShipState

	level    int
	capacity int

	ressources map[string]int
	money      float32

	docked *Harbour

	danger Agent

	behaviourPercept    func(*Ship, *Environment, int)
	behaviourDeliberate func(*Ship, *Environment, int)
	behaviourAct        func(*Ship, *Environment, int)
	behaviourLevelUp    func(*Ship, *Environment, int)
}

// Different types of ships
// Pirate, Merchant, Privateer, Army, Filibusteer, PirateHunter, Zoro

// Pirate ship
func NewShipPirate(id string, coor Coordinates, docked *Harbour) *Ship {
	shipType := "Pirate"
	config := constants.ShipConfigs[shipType]
	ressources := make(map[string]int)
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Ship",
			fov:        config.FOV,
			speed:      config.Speed,
			attack:     config.Attack,
			health:     config.Health,
		},
		shipType:    shipType,
		nationality: constants.Without,
		level:       1,
		capacity:    config.Capacity,
		ressources:  ressources,

		docked: docked,

		behaviourPercept:    pirateShipPercept,
		behaviourDeliberate: pirateShipDeliberate,
		behaviourAct:        pirateShipAct,
		behaviourLevelUp:    pirateShipLevelUp,
	}
}

// Merchant ship
func NewShipMerchant(id string, nationality constants.Nationality, coor Coordinates, docked *Harbour) *Ship {
	shipType := "Merchant"
	config := constants.ShipConfigs[shipType]
	ressources := make(map[string]int)
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Ship",
			fov:        config.FOV,
			speed:      config.Speed,
			attack:     config.Attack,
			health:     config.Health,
		},
		shipType:    shipType,
		nationality: nationality,
		level:       1,
		ressources:  ressources,
		capacity:    config.Capacity,

		docked: docked,

		behaviourPercept:    merchantShipPercept,
		behaviourDeliberate: merchantShipDeliberate,
		behaviourAct:        merchantShipAct,
		behaviourLevelUp:    merchantShipLevelUp,
	}
}

// Privateer ship
func NewShipPrivateer(id string, nationality constants.Nationality, coor Coordinates, docked *Harbour) *Ship {
	shipType := "Privateer"
	config := constants.ShipConfigs[shipType]
	ressources := make(map[string]int)
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Ship",
			fov:        config.FOV,
			speed:      config.Speed,
			attack:     config.Attack,
			health:     config.Health,
		},
		shipType:    shipType,
		nationality: nationality,
		level:       1,
		ressources:  ressources,
		capacity:    config.Capacity,

		docked: docked,

		behaviourPercept:    privateerShipPercept,
		behaviourDeliberate: privateerShipDeliberate,
		behaviourAct:        privateerShipAct,
		behaviourLevelUp:    privateerShipLevelUp,
	}
}

// Army ship
func NewShipArmy(id string, nationality constants.Nationality, coor Coordinates, docked *Harbour) *Ship {
	shipType := "Army"
	config := constants.ShipConfigs[shipType]
	ressources := make(map[string]int)
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Ship",
			fov:        config.FOV,
			speed:      config.Speed,
			attack:     config.Attack,
			health:     config.Health,
		},
		shipType:    shipType,
		nationality: nationality,
		level:       1,
		ressources:  ressources,

		capacity: config.Capacity,

		docked: docked,

		behaviourPercept:    armyShipPercept,
		behaviourDeliberate: armyShipDeliberate,
		behaviourAct:        armyShipAct,
		behaviourLevelUp:    armyShipLevelUp,
	}
}

// Zoro ship
func NewShipZoro(id string, coor Coordinates, docked *Harbour) *Ship {
	shipType := "Zoro"
	config := constants.ShipConfigs[shipType]
	ressources := make(map[string]int)
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Ship",
			fov:        config.FOV,
			speed:      config.Speed,
			attack:     config.Attack,
			health:     config.Health,
		},
		shipType:    shipType,
		nationality: constants.Without,
		level:       1,
		ressources:  ressources,
		capacity:    config.Capacity,

		docked: docked,

		behaviourPercept:    zoroShipPercept,
		behaviourDeliberate: zoroShipDeliberate,
		behaviourAct:        zoroShipAct,
		behaviourLevelUp:    zoroShipLevelUp,
	}
}

// Rackham ship
func NewShipRackham(id string, coor Coordinates, docked *Harbour) *Ship {
	shipType := "Rackham"
	config := constants.ShipConfigs[shipType]
	ressources := make(map[string]int)
	return &Ship{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Ship",
			fov:        config.FOV,
			speed:      config.Speed,
			attack:     config.Attack,
			health:     config.Health,
		},
		shipType:    shipType,
		nationality: constants.Without,
		level:       1,
		ressources:  ressources,
		capacity:    config.Capacity,

		docked: docked,

		behaviourPercept:    rackhamShipPercept,
		behaviourDeliberate: rackhamShipDeliberate,
		behaviourAct:        rackhamShipAct,
		behaviourLevelUp:    rackhamShipLevelUp,
	}
}

// Getter Nationality
func (ag *Ship) Nationality() constants.Nationality {
	return ag.nationality
}

// Setter Nationality
func (ag *Ship) SetNationality(nationality constants.Nationality) {
	ag.Lock()
	defer ag.Unlock()

	ag.nationality = nationality
}

// Getter ShipState
func (ag *Ship) ShipState() ShipState {
	return ag.shipState
}

// Setter ShipState
func (ag *Ship) SetShipState(state ShipState) {
	ag.shipState = state
}

// Decide current ShipState
func (ag *Ship) CurrentShipState() {
	if ag.Action() == "attack" {
		ag.SetShipState(Attacking)
		return
	}

	if ag.Health() <= 20 {
		ag.SetShipState(OnFire)
		return
	}

	ag.SetShipState(Idle)
}

// Getter ShipType
func (ag *Ship) ShipType() string {
	return ag.shipType
}

// Getter Level
func (ag *Ship) Level() int {
	return ag.level
}

func (ag *Ship) LevelUp(env *Environment, step int) {
	ag.behaviourLevelUp(ag, env, step)
}

// Getter Capacity
func (ag *Ship) Capacity() int {
	ag.RLock()
	defer ag.RUnlock()

	return ag.capacity
}

// Setter Capacity
func (ag *Ship) SetCapacity(capacity int) {
	ag.Lock()
	defer ag.Unlock()

	ag.capacity = capacity
}

// Getter Ressources
func (ag *Ship) Ressources() map[string]int {
	ag.RLock()
	defer ag.RUnlock()

	return ag.ressources
}

// Setter Ressources
func (ag *Ship) SetRessources(ressources map[string]int) {
	ag.Lock()
	defer ag.Unlock()

	ag.ressources = ressources
}

// Setter Ressources Add
func (ag *Ship) AddRessources(ressource string, quantity int) {
	ag.Lock()
	defer ag.Unlock()

	ag.ressources[ressource] += quantity
}

// Setter Ressources Remove
func (ag *Ship) RemoveRessources(ressource string, quantity int) {
	ag.Lock()
	defer ag.Unlock()

	currentQuantity, exists := ag.ressources[ressource]

	// If it does not exist do nothing
	if !exists {
		return
	}

	// If it uses all the quantity remove totally
	if quantity >= currentQuantity {
		ag.ressources[ressource] = 0
	} else {
		ag.ressources[ressource] -= quantity
	}
}

// Getter Money
func (ag *Ship) Money() float32 {
	ag.RLock()
	defer ag.RUnlock()

	return ag.money
}

// Setter Money
func (ag *Ship) SetMoney(money float32) {
	ag.Lock()
	defer ag.Unlock()

	ag.money = money
}

// Setter Money Add
func (ag *Ship) AddMoney(money float32) {
	ag.Lock()
	defer ag.Unlock()

	ag.money = ag.money + money
}

// Setter Money Remove
func (ag *Ship) RemoveMoney(money float32) {
	ag.Lock()
	defer ag.Unlock()

	ag.money = ag.money - money
}

// Getter Docked
func (ag *Ship) Docked() *Harbour {
	ag.RLock()
	defer ag.RUnlock()

	return ag.docked
}

// Setter Docked
func (ag *Ship) SetDocked(harbour *Harbour) {
	ag.Lock()
	defer ag.Unlock()

	ag.docked = harbour
}

// Getter Danger
func (ag *Ship) Danger() Agent {
	return ag.danger
}

// Setter Danger
func (ag *Ship) SetDanger(danger Agent) {
	ag.danger = danger
}

func (ag *Ship) Percept(env *Environment, step int) {
	ag.behaviourPercept(ag, env, step)
}

func (ag *Ship) Deliberate(env *Environment, step int) {
	ag.behaviourDeliberate(ag, env, step)
}

func (ag *Ship) Act(env *Environment, step int) {
	ag.behaviourAct(ag, env, step)
}

// Repair the ship for 1 health is 1 money to the max of health or money
func (ag *Ship) Repair() {
	repairsNeeded := constants.ShipConfigs[ag.ShipType()].Health - ag.Health()
	repairsNeeded = min(repairsNeeded, int(ag.Money()))
	ag.AddHealth(repairsNeeded)
	ag.RemoveMoney(float32(repairsNeeded))
}

func (ag *Ship) Quantity() int {
	totalResources := int(0)

	// Calculate total resources on the ship
	for _, quantity := range ag.Ressources() {
		totalResources += quantity
	}

	return totalResources
}

/*
--------------------------------------
*********** Kraken Agent *************
--------------------------------------
*/
// The Kraken finds closest agents and attacks them

type Kraken struct {
	*BaseAgent

	sync.RWMutex

	krakenState KrakenState
}

func NewKraken(id string, coor Coordinates) *Kraken {
	return &Kraken{
		BaseAgent: &BaseAgent{
			BaseObject: BaseObject{coordinates: coor},
			id:         id,
			agtType:    "Kraken",
			fov:        constants.KrakenParameters.FOV,
			speed:      constants.KrakenParameters.Speed,
			attack:     constants.KrakenParameters.Attack,
			health:     constants.KrakenParameters.Health,
		},
	}
}

// Getter KrakenState
func (ag *Kraken) KrakenState() KrakenState {
	return ag.krakenState
}

// Setter KrakenState
func (ag *Kraken) SetKrakenState(krakenState KrakenState) {
	ag.krakenState = krakenState
}

// Compute current KrakenState
func (ag *Kraken) CurrentKrakenState() {
	if ag.Action() == "rest" {
		ag.SetKrakenState(Hidden)
		return
	}

	ag.SetKrakenState(OutOfWater)
}

// Regain 10% of max health per step resting
func (ag *Kraken) RestAndGainHealth() {
	ag.AddHealth(int(float32(constants.KrakenParameters.Health) * 0.05))
}

func (ag *Kraken) Percept(env *Environment, step int) {
	// Retrieve the closest agent from kraken
	closest := env.ClosestShipFromKraken(ag)

	// Check if there is a closest ship within fov
	if closest != nil {
		ag.SetTarget(closest)
	} else {
		ag.SetTarget(nil)
	}
}

func (ag *Kraken) Deliberate(env *Environment, step int) {
	// If the Kraken is low life or has no target = rest to gain PV
	if ag.Health() <= int(float32(constants.KrakenParameters.Health)*0.05) {
		ag.SetAction("rest")
		return
	}

	if ag.Target() == nil && len(ag.Itinerary()) == 0 {
		destination := NewRandomCoordinates()
		path := AStar(ag.Coordinates(), destination)
		if path != nil {
			ag.SetItinerary(path)
		}
		ag.SetAction("follow")
		return
	}

	//If not finsihed itinerary just follow
	if ag.Target() == nil && len(ag.Itinerary()) > 0 {
		ag.SetAction("follow")
		return
	}

	// Calculate the distance to the target agent (Manhattan)
	distance := Heuristic(ag.Coordinates(), (ag.Target()).Coordinates())

	// Ensure the target is a ship and not docked
	ship, ok := (ag.Target()).(*Ship)
	if ok && ship.Docked() != nil {
		ag.SetAction("rest")
		return
	}

	// If the Kraken can attack, it should attack the ship
	if distance < ag.Speed() {
		ag.SetAction("attack")
	} else {
		// Pathfinding to the target
		path := AStar(ag.Coordinates(), (ag.Target()).Coordinates())
		// If path found set itinerary and follow
		if path != nil {
			ag.SetItinerary(path)
			ag.SetAction("follow")
		}
	}
}

func (ag *Kraken) Act(env *Environment, step int) {

	//5% chance of mistake
	rand.Seed(uint64(time.Now().UnixNano()))
	mistake := rand.Intn(20)

	if mistake == 0 {
		ag.SetAction("rest")
	}

	switch ag.Action() {
	case "attack":
		ship, ok := (ag.Target()).(*Ship)
		if ok && ship.Docked() != nil {
			break
		}
		// Go to target and attack
		ag.FollowItinerary(true)
		env.AttackAgent(ag.Target(), ag)
	case "follow":
		// Move towards target
		ag.FollowItinerary(false)
	case "rest":
		ag.RestAndGainHealth()
	default:
		ag.RestAndGainHealth()
	}

	ag.target = nil
	ag.action = ""

	//Compute KrakenState and Direction for sprite
	ag.CurrentKrakenState()
	ag.CurrentDirection()
}
