package UI

import (
	"sync"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/controllers"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

type Game struct {
	ebiten.Game
	Ui         *ebitenui.UI
	Controller controllers.SimulationController
	Simu       *simulation.Simulation

	AgentLock sync.RWMutex
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	g.Ui.Update()
	return nil
}
