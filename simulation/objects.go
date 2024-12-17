package simulation

import "sync"

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
	sync.RWMutex
	coordinates Coordinates
}

// Getter coordinates
func (o *BaseObject) Coordinates() Coordinates {
	o.RLock()
	defer o.RUnlock()
	return o.coordinates
}

// Setter coordinates
func (o *BaseObject) SetCoordinates(coor Coordinates) {
	o.Lock()
	defer o.Unlock()
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
	*BaseObject
	name                string
	sellableRessources  map[string]bool
	availableRessources map[string]bool
}

func NewHarbour(name string, coor Coordinates) *Harbour {
	sellableRessources := make(map[string]bool)
	availableRessources := make(map[string]bool)
	return &Harbour{
		name:                name,
		BaseObject:          &BaseObject{coordinates: coor},
		sellableRessources:  sellableRessources,
		availableRessources: availableRessources,
	}
}

// Getter
func (h *Harbour) Name() string {
	return h.name
}

// Getter
func (h *Harbour) SellableRessources() map[string]bool {
	return h.sellableRessources
}

// Getter
func (h *Harbour) AvailableRessources() map[string]bool {
	return h.availableRessources
}

// Add sellable ressource
func (h *Harbour) AddSellableRessources(ressource string) {
	h.sellableRessources[ressource] = true
}

// Add available ressource
func (h *Harbour) AddAvailableRessources(ressource string) {
	h.availableRessources[ressource] = true
}

// IsRessourcesSellable checks if a ressource is sellable in a Harbour
func (h *Harbour) IsRessourcesSellable(ressource string) bool {

	// Get all sellable ressources
	sellableRessources := h.SellableRessources()

	// Check if the ressource given by agent is sellable
	sellable, ok := sellableRessources[ressource]
	if ok && sellable {
		return true
	}
	return false
}
