package simulation

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Pirateship struct {
	sprite.Sprite
}

func NewPirateship(position engine.Position) *Pirateship {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, engine.ShipImage)
	if err != nil {
		fmt.Printf("création de PirateShip impossible : %s", err)
	}

	return &Pirateship{*sprite}
}

func (p *Pirateship) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *Pirateship) ToString() string {
	return "this is the pirateship!"
}
