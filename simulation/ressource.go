package simulation

/*
---------------------------------------
************* Ressource ***************
---------------------------------------
*/

type Ressource struct {
	name            string
	initPrice       float32
	minPrice        float32
	hyperbolicDecay float32
}
