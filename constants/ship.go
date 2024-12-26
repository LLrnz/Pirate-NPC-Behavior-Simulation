package constants

// ShipType defines types of ships in the game
type ShipType = string

const (
	Pirate    ShipType = "Pirate"
	Merchant  ShipType = "Merchant"
	Privateer ShipType = "Privateer"
	Army      ShipType = "Army"
	Zoro      ShipType = "Zoro"
	Rackham   ShipType = "Rackham"
)

// ShipConfig is base config for ship creation
type ShipConfig struct {
	Speed    int
	Attack   int
	Health   int
	Capacity int
	FOV      int
}

// ShipConfigs has configurations of ships
var ShipConfigs = map[ShipType]ShipConfig{
	Pirate: {
		Speed:    1,
		Attack:   50,
		Health:   200,
		Capacity: 10,
		FOV:      10,
	},
	Merchant: {
		Speed:    1,
		Attack:   5,
		Health:   300,
		Capacity: 50,
		FOV:      12,
	},
	Privateer: {
		Speed:    2,
		Attack:   60,
		Health:   300,
		Capacity: 20,
		FOV:      12,
	},
	Army: {
		Speed:    1,
		Attack:   100,
		Health:   300,
		Capacity: 15,
		FOV:      10,
	},
	Zoro: {
		Speed:    1,
		Attack:   400,
		Health:   400,
		Capacity: 5,
		FOV:      25,
	},
	Rackham: {
		Speed:    1,
		Attack:   5,
		Health:   400,
		Capacity: 5,
		FOV:      10,
	},
}
