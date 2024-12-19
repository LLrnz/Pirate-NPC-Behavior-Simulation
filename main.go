package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
	"gitlab.utc.fr/arsaintm/ia04-projet/view"
)

func main() {
	// Init the Model
	simu := simulation.NewSimulation(600*time.Second, 0)

	// Init the View
	game := view.NewGame(simu, 1536, 864)

	// Run the simulation in a separate goroutine so it works at the same time as the game
	go simu.Run()

	// Start the Ebitengine game
	ebiten.SetWindowSize(1024, 576)
	ebiten.SetWindowTitle("Pirateeees")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
