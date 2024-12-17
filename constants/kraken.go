package constants

// KrakenConfig
type KrakenConfig struct {
	Speed  int
	Attack int
	Health int
	FOV    int
}

// Kraken config for the simulation
var KrakenParameters = KrakenConfig{
	Speed:  1,
	Attack: 200,
	Health: 1000,
	FOV:    10,
}
