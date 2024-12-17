package utils

const (
	//Size of the background
	MapWidth = 768
	MapHeigt = 432

	Zoom = 1

	//size of the simulation view
	GameWidth  = MapWidth * Zoom
	GameHeight = MapHeigt * Zoom

	WindowWidth     = 868
	WindowHeight    = 432
	MaxWindowWidth  = 968
	MaxWindowHeight = 532

	//Size of a tile on screen
	TileSize = 16 * Zoom

	//Asset of the background
	MapImage = "assets/mapMonde.png"

	//Kraken sprites
	KrakenIdle   = "assets/kraken.png"
	KrakenHidden = "assets/kraken_shadow.png"

	//army sprites
	ArmyIdle   = "assets/armyship.png"
	ArmyAttack = "assets/armyship_attack.png"
	ArmyFire   = "assets/armyship_fire.png"

	//pirate sprites
	PirateIdle   = "assets/pirateship.png"
	PirateAttack = "assets/pirateship_attack.png"
	PirateFire   = "assets/pirateship_fire.png"

	//privateer sprites
	PrivateerIdle   = "assets/huntership.png"
	PrivateerAttack = "assets/huntership_attack.png"
	PrivateerFire   = "assets/huntership_fire.png"

	//merchant sprites
	MerchantIdle   = "assets/merchantship.png"
	MerchantAttack = "assets/merchantship_attack.png"
	MerchantFire   = "assets/merchantship_fire.png"

	//zoro sprites
	ZoroIdle   = "assets/zoroship.png"
	ZoroAttack = "assets/zoroship_attack.png"
	ZoroFire   = "assets/zoroship_fire.png"
)
