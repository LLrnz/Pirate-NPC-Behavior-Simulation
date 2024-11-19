package main

import (
	"image"
	"log"
	"time"

	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
	"simulationPirate/pkg/engine/ui"
	"simulationPirate/pkg/simulation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	fond   *sprite.Sprite
	ticker *time.Ticker
}

func (g *Game) LoadContent() {
	var err error
	g.fond, err = sprite.NewSprite(engine.Position{Point: image.Point{0, 0}}, sprite.Right, 4, engine.MapImage)
	if err != nil {
		log.Fatal(err)
	}
	gridSize := grid.ToCoor(engine.Position{Point: image.Point{engine.ScreenHeight, engine.ScreenWidth}})
	grid.GetInstance().Init(gridSize.X, gridSize.Y)

	shipPos := engine.Position{Point: image.Point{0, 0}}
	grid.GetInstance().SetObjectAtPos(shipPos, simulation.NewPirateship(shipPos))

	g.ticker = time.NewTicker(time.Second * 2)

}

func (g *Game) Update() error {
	go func() {
		select {
		case <-g.ticker.C:
			grid.GetInstance().Update()

			break
		}
	}()
	ui.Update()
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
