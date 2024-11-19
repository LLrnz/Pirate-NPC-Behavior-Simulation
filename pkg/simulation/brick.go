package simulation

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Brick struct {
	sprite.Sprite
}

func NewBrick(position image.Point) *Brick {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Brick impossible : %s", err)
	}

	return &Brick{*sprite}
}

func (p *Brick) Update() {
	next := grid.ToPos(grid.GetInstance().GetNeighbor(p.Position, p.Direction))
	p.Position = next
	fmt.Print(next)
}
