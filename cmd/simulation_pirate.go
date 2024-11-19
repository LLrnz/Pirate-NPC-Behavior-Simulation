package main

import (
	"image"
	"log"
	"time"

	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
	"simulationPirate/pkg/simulation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	fond *sprite.Sprite
}

func (g *Game) LoadContent() {
	var err error
	g.fond, err = sprite.NewSprite(image.Point{0, 0}, sprite.Right, 4, "assets/mapMonde.png")
	if err != nil {
		log.Fatal(err)
	}
	grid.GetInstance().Init(engine.ScreenHeight, engine.ScreenWidth)

	shipPos := image.Point{0, 0}
	grid.GetInstance().Grid[shipPos] = simulation.NewBrigantin(shipPos)

}

func (g *Game) Update() error {
	grid.GetInstance().Update()
	time.Sleep(time.Duration(2 * time.Second))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.fond.Draw(screen)

	grid.GetInstance().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 256, 144
}

func main() {
	ebiten.SetWindowSize(1024, 576)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}
	g.LoadContent()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
