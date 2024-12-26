package simulation

// Ressource struct
type Ressource struct {
	name            string
	initPrice       float32
	minPrice        float32
	hyperbolicDecay float32
}

var Gold = Ressource{"Gold", 110.0, 5.0, 6.0}
var Silver = Ressource{"Silver", 105.0, 5.0, 7.0}
var Coffee = Ressource{"Coffee", 100.0, 5.0, 8.0}
var Spices = Ressource{"Spices", 95.0, 5.0, 9.0}
var Sugar = Ressource{"Sugar", 90.0, 5.0, 10.0}
var Rum = Ressource{"Rum", 85.0, 5.0, 11.0}
var Tobacco = Ressource{"Tobacco", 80.0, 5.0, 12.0}
var Silk = Ressource{"Silk", 75.0, 5.0, 13.0}
var Cocoa = Ressource{"Cocoa", 70.0, 5.0, 14.0}
var Pearls = Ressource{"Pearls", 65.0, 5.0, 15.0}
var Cotton = Ressource{"Cotton", 60.0, 5.0, 16.0}
var Ivory = Ressource{"Ivory", 55.0, 5.0, 17.0}
var Weapons = Ressource{"Weapons", 50.0, 5.0, 18.0}
var Tipiak = Ressource{"Tipiak", 45.0, 5.0, 19.0}

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
