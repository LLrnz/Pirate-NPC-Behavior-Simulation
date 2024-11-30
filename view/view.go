package view

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

/*
----------------------------------------
***************** View *****************
----------------------------------------
*/
// View handles the display of the game separately from the Model

// Game struct is used by ebiten to show the window
type Game struct {
	simu      *simulation.Simulation
	agentLock sync.RWMutex
	width     int
	height    int
}

// NewGame creates a new Game
func NewGame(simu *simulation.Simulation, width, height int) *Game {
	return &Game{
		simu:   simu,
		width:  width,
		height: height,
	}
}

// Update [[WRITE HERE WHAT IT DOES]]
func (g *Game) Update() error {
	// [[WE WILL NEED TO UPDATE HERE THE INTERACTIVE ELEMENTS]]
	return nil
}

// Draw takes infos from the Model and draw it on the window
func (g *Game) Draw(screen *ebiten.Image) {
	g.agentLock.RLock()
	defer g.agentLock.RUnlock()

	// Clear the screen with a background color.
	img, err := getImageFromFilePath("assets/mapMonde.png")
	if err != nil {
		fmt.Println("Error loading image:", err)
		return // or handle gracefully
	}
	screen.DrawImage(img, nil)

	// Size of the squares representing agents.
	const agentSize = 10.0

	// Draw each agent as a rectangle.
	for _, agent := range g.simu.Agents() {
		coor := agent.Coordinates()
		x := float32(coor.X())
		y := float32(coor.Y())

		// Choose color based on agent type.
		var clr color.RGBA
		switch agent.AgtType() {
		case "ship":
			clr = color.RGBA{0, 255, 0, 255} // Green for ships
		case "kraken":
			clr = color.RGBA{255, 0, 0, 255} // Red for kraken
		default:
			clr = color.RGBA{255, 255, 255, 255} // White for others
		}

		// Draw the filled rectangle.
		vector.DrawFilledRect(screen, x, y, agentSize, agentSize, clr, false)
	}
}

// Layout sets the window dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
