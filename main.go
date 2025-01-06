package main

import (
	"log"
	"sync"
	"time"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

func main() {
	// Init the Model
	simu := simulation.NewSimulation(time.Hour, 3654)
	mw, controller := UI.NewMainWindow(simu, 30)

	// Ebiten setup
	ebiten.SetWindowSize(utils.WindowWidth, utils.WindowHeight)
	ebiten.SetWindowSizeLimits(utils.MaxWindowWidth, utils.MaxWindowHeight, -1, -1)
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

	game.Controller.Start()

	// Run the simulation in a separate goroutine so it works at the same time as the game
	go simu.Run()

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}

}
