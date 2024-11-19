package simulation

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Pirateship struct {
	sprite.Sprite
}

func NewPirateship(position image.Point) *Pirateship {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de PirateShip impossible : %s", err)
	}

	return &Pirateship{*sprite}
}

func (p *Pirateship) Update() {
	next := grid.ToPos(grid.GetInstance().GetNeighbor(p.Position, p.Direction))
	p.Position = next
	fmt.Print(next)
}
