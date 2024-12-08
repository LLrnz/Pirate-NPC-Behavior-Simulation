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
		Attack:   20,
		Health:   100,
		Capacity: 10,
		FOV:      10,
	},
	Merchant: {
		Speed:    2,
		Attack:   5,
		Health:   150,
		Capacity: 50,
		FOV:      12,
	},
	Privateer: {
		Speed:    3,
		Attack:   15,
		Health:   120,
		Capacity: 20,
		FOV:      14,
	},
	Army: {
		Speed:    1,
		Attack:   25,
		Health:   200,
		Capacity: 15,
		FOV:      8,
	},
	Zoro: {
		Speed:    4,
		Attack:   40,
		Health:   200,
		Capacity: 5,
		FOV:      20,
	},
	Rackham: {
		Speed:    5,
		Attack:   40,
		Health:   200,
		Capacity: 5,
		FOV:      20,
	},
}
