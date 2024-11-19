package simulation

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type TroisMatesCarre struct {
	sprite.Sprite
}

func NewTroisMatesCarre(position image.Point) *TroisMatesCarre {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de TroisMatesCarre impossible : %s", err)
		return nil
	}

	return &TroisMatesCarre{*sprite}
}

func (p *TroisMatesCarre) Update() {
	next := grid.ToPos(grid.GetInstance().GetNeighbor(p.Position, p.Direction))
	p.Position = next
	fmt.Print(next)
}
