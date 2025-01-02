package utils

const (
	//Size of the background
	MapWidth = 768 * 2
	MapHeigt = 432 * 2

	Zoom = 1

	//size of the simulation view
	GameWidth  = MapWidth * Zoom
	GameHeight = MapHeigt * Zoom

	WindowWidth     = MapWidth + 100
	WindowHeight    = MapHeigt
	MaxWindowWidth  = MapWidth + 200
	MaxWindowHeight = +100

	//Size of a tile on screen
	TileSize = 16 * Zoom

	//Asset of the background
	MapImage = "assets/worldMap_big_2.png"

	//TODO : add all the asset path ?
)
