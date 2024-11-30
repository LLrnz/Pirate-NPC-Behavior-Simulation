package simulation

/*
********************************************
	objects.go - Objects logique
********************************************
*/

/*
----------------------------------------
******* Coordinates of objects *********
----------------------------------------
*/
// Coordinates are used by the map to position objects/agents

type Coordinates struct {
	x int
	y int
}

// Getter X coord
func (c *Coordinates) X() int {
	return c.x
}

// Getter Y coord
func (c *Coordinates) Y() int {
	return c.y
}

// Setter X coord
func (c *Coordinates) SetX(x int) {
	c.x = x
}

// Setter Y coord
func (c *Coordinates) SetY(y int) {
	c.y = y
}

/*
----------------------------------------
********** Object Interface ************
----------------------------------------
*/

type Object interface {
	Coordinates() Coordinates
	SetCoordinates(coor Coordinates)
}

/*
----------------------------------------
************* BaseObject ***************
----------------------------------------
*/
// It is the object struct that is used by any object

// Creation of BaseObject so it can be used by other structs
type BaseObject struct {
	coordinates Coordinates
}

// Getter coordinates
func (o *BaseObject) Coordinates() Coordinates {
	return o.coordinates
}

// Setter coordinates
func (o *BaseObject) SetCoordinates(coor Coordinates) {
	o.coordinates = coor
}

/*
---------------------------------------
************** Harbour ****************
---------------------------------------
*/
// The Harbour is an object where Ships can dock to be safe
// While docked, they can repair or upgrade their ship
// Merchants ships can also take ressources or sell ressources for money

// Creation of our Harbour object that takes a BaseObject
type Harbour struct {
	BaseObject
	name                string
	sellableRessources  map[string]bool
	availableRessources map[string]bool
	dockedAgents        []Agent
}

func NewHarbour(name string, coor Coordinates) *Harbour {
	return &Harbour{
		name:       name,
		BaseObject: BaseObject{coordinates: coor},
	}
}
