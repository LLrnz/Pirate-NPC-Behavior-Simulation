package simulation

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Brigantin struct {
	sprite.Sprite
}

func NewBrigantin(position image.Point) *Brigantin {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Brigantin impossible : %s", err)
		return nil
	}

	return &Brigantin{*sprite}
}

func (p *Brigantin) Update() {
	next := grid.ToPos(grid.GetInstance().GetNeighbor(p.Position, p.Direction))
	p.Position = next
	fmt.Print(next)
}
