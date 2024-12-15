package simulation

import "gitlab.utc.fr/arsaintm/ia04-projet/constants"

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

	Target() *Agent
	SetTarget(target *Agent)
	Action() string
	SetAction(action string)
	Itinerary() []Coordinates
	SetItinerary(itinerary []Coordinates)

	FollowItinerary()

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
	id      string
	agtType string

	fov    int
	speed  int
	attack int
	health int

	target    *Agent
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
	return agt.attack
}

// Setter Attack
func (agt *BaseAgent) SetAttack(attack int) {
	agt.attack = attack
}

// Getter Health
func (agt *BaseAgent) Health() int {
	return agt.health
}

// Setter Health
func (agt *BaseAgent) SetHealth(health int) {
	agt.health = health
}

// Setter Health Add
func (agt *BaseAgent) AddHealth(health int) {
	agt.health = agt.health + health
}

// Setter Health Remove
func (agt *BaseAgent) RemoveHealth(health int) {
	agt.health = agt.health - health
}

// Getter Target
func (agt *BaseAgent) Target() *Agent {
	return agt.target
}

// Setter Target
func (agt *BaseAgent) SetTarget(target *Agent) {
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
	if agt.Itinerary() == nil {
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
func (ag *BaseAgent) FollowItinerary() {
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
}

/*
--------------------------------------
*************** Ship *****************
--------------------------------------
*/
// There are different types of ships [[Explain here what behaviours they will have]]

type Ship struct {
	*BaseAgent
	shipType    string
	nationality constants.Nationality
	shipState   ShipState

	level    int
	capacity int

	ressources map[string]int
	money      float32

	docked *Harbour

	danger *Agent

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
	return ag.capacity
}

// Setter Capacity
func (ag *Ship) SetCapacity(capacity int) {
	ag.capacity = capacity
}

// Getter Ressources
func (ag *Ship) Ressources() map[string]int {
	return ag.ressources
}

// Setter Ressources
func (ag *Ship) SetRessources(ressources map[string]int) {
	ag.ressources = ressources
}

// Setter Ressources Add
func (ag *Ship) AddRessources(ressource string, quantity int) {
	ag.ressources[ressource] += quantity
}

// Setter Ressources Remove
func (ag *Ship) RemoveRessources(ressource string, quantity int) {
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
	return ag.money
}

// Setter Money
func (ag *Ship) SetMoney(money float32) {
	ag.money = money
}

// Setter Money Add
func (ag *Ship) AddMoney(money float32) {
	ag.money = ag.money + money
}

// Setter Money Remove
func (ag *Ship) RemoveMoney(money float32) {
	ag.money = ag.money - money
}

// Getter Docked
func (ag *Ship) Docked() *Harbour {
	return ag.docked
}

// Setter Docked
func (ag *Ship) SetDocked(harbour *Harbour) {
	ag.docked = harbour
}

// Getter Danger
func (ag *Ship) Danger() *Agent {
	return ag.danger
}

// Setter Danger
func (ag *Ship) SetDanger(danger *Agent) {
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
	ag.AddHealth(int(float32(constants.KrakenParameters.Health) * 0.1))
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

	//If the Kraken is low life or has no target = rest to gain PV
	if ag.Health() <= int(float32(constants.KrakenParameters.Health)*0.1) || ag.Target() == nil {
		ag.SetAction("rest")
		return
	}

	// Calculate the distance to the target agent (Manhattan)
	distance := Heuristic(ag.Coordinates(), (*ag.Target()).Coordinates())

	// If the Kraken can attack, it should attack the ship
	// Else it should follow the nearest agent as a target
	// Skip docked ships because they are safe
	ship, ok := (*ag.Target()).(*Ship)

	if (distance < ag.Speed() && ok && ship.Docked() == nil) || (distance < ag.Speed() && !ok) {
		ag.SetAction("attack")
	} else {
		// Pathfinding to the target
		path := AStar(ag.Coordinates(), (*ag.Target()).Coordinates())
		// If path found set itinerary and follow
		if path != nil {
			ag.SetItinerary(path)
			ag.SetAction("follow")
		}
	}

}

func (ag *Kraken) Act(env *Environment, step int) {

	switch ag.Action() {
	case "attack":
		// Go to target and attack
		ag.FollowItinerary()
		// Remove the target from the environment
		env.AttackAgent(*ag.Target(), ag)
	case "follow":
		// Move towards target
		ag.FollowItinerary()
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
