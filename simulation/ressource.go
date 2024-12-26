package simulation

// Ressource struct
type Ressource struct {
	name            string
	initPrice       float32
	minPrice        float32
	hyperbolicDecay float32
}

var Gold = Ressource{"Gold", 600.0, 40.0, 4.0}
var Silver = Ressource{"Silver", 450.0, 300.0, 12.0}
var Coffee = Ressource{"Coffee", 120.0, 90.0, 1.5}
var Spices = Ressource{"Spices", 300.0, 200.0, 2.5}
var Sugar = Ressource{"Sugar", 150.0, 100.0, 2.0}
var Rum = Ressource{"Rum", 50.0, 30.0, 1.0}
var Tobacco = Ressource{"Tobacco", 80.0, 60.0, 1.2}
var Silk = Ressource{"Silk", 500.0, 350.0, 18.0}
var Cocoa = Ressource{"Cocoa", 180.0, 140.0, 2.0}
var Pearls = Ressource{"Pearls", 700.0, 500.0, 40.0}
var Cotton = Ressource{"Cotton", 50.0, 35.0, 0.8}
var Ivory = Ressource{"Ivory", 400.0, 250.0, 60.5}
var Weapons = Ressource{"Weapons", 250.0, 180.0, 2.5}
var Tipiak = Ressource{"Tipiak", 300.0, 200.0, 3.0}

// Ressources
var RessourceMap = map[string]Ressource{
	"Gold":    Gold,
	"Silver":  Silver,
	"Coffee":  Coffee,
	"Spices":  Spices,
	"Sugar":   Sugar,
	"Rum":     Rum,
	"Tobacco": Tobacco,
	"Silk":    Silk,
	"Cocoa":   Cocoa,
	"Pearls":  Pearls,
	"Cotton":  Cotton,
	"Ivory":   Ivory,
	"Weapons": Weapons,
	"Tipiak":  Tipiak,
}

// GetRessource retrieves a Ressource by its name
func GetRessource(name string) (Ressource, bool) {
	ressource, found := RessourceMap[name]
	return ressource, found
}
