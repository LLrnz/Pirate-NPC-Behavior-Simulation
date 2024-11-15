package main

import (
	"image"
	"log"

	"simulationPirate/pkg/sprite"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	fond *sprite.Sprite
}

func (g *Game) LoadContent() {
	var err error
	g.fond, err = sprite.NewSprite(&image.Point{0, 0}, sprite.Right, 4, "assets/mapMonde.png")

	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.fond.Draw(screen)
	ebitenutil.DebugPrint(screen, "Hello, World!")
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
