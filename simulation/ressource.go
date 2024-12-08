package simulation

// Ressource struct
type Ressource struct {
	name            string
	initPrice       float32
	minPrice        float32
	hyperbolicDecay float32
}

// Resources
var (
	Gold    = Ressource{"Gold", 100.0, 50.0, 0.01}
	Coffee  = Ressource{"Coffee", 20.0, 10.0, 0.02}
	Spices  = Ressource{"Spices", 25.0, 12.0, 0.015}
	Sugar   = Ressource{"Sugar", 15.0, 8.0, 0.018}
	Alcohol = Ressource{"Alcohol", 30.0, 15.0, 0.025}
	Tipiak  = Ressource{"Tipiak", 50.0, 25.0, 0.03}
)

// Ressources
var RessourceMap = map[string]Ressource{
	"Gold":    Gold,
	"Coffee":  Coffee,
	"Spices":  Spices,
	"Sugar":   Sugar,
	"Alcohol": Alcohol,
	"Tipiak":  Tipiak,
}

// GetRessource retrieves a Ressource by its name
func GetRessource(name string) (Ressource, bool) {
	ressource, found := RessourceMap[name]
	return ressource, found
}
