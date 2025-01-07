package main

import (
	"bytes"
	"log"
	"os"
	"sync"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

func PlayMP3() {
	file, _ := os.Open("assets/audio/pirates.mp3")
	defer file.Close()

	decoder, _ := mp3.NewDecoder(file)

	sampleRate := decoder.SampleRate()

	ctx, _ := oto.NewContext(sampleRate, 2, 2, 8192)

	player := ctx.NewPlayer()
	defer player.Close()

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(decoder)

	_, _ = player.Write(buf.Bytes())
}

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

	go PlayMP3()

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}

}
