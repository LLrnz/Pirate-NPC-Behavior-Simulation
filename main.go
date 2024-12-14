package main

import (
	"log"
	"sync"
	"time"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

func main() {
	// Init the Model
	simu := simulation.NewSimulation(600*time.Second, 0)
	mw, controller := UI.NewMainWindow(simu, 0.5)

	// Init the View

	// Ebiten setup
	ebiten.SetWindowSize(868, 432)
	ebiten.SetWindowSizeLimits(968, 532, -1, -1)
	ebiten.SetWindowTitle("Piraaaates")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ui := ebitenui.UI{
		Container: mw,
	}

	game := UI.Game{
		Ui:         &ui,
		Controller: controller,
		Simu:       simu,
		AgentLock:  sync.RWMutex{}}

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}

	// Run the simulation in a separate goroutine so it works at the same time as the game
	go simu.Run()

}
