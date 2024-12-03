package models

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Kraken struct {
	sprite.Sprite
}

func NewKraken(position engine.Position) *Kraken {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Kraken impossible : %s", err)
		return nil
	}

	return &Kraken{*sprite}
}

func (p *Kraken) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *Kraken) ToString() string {
	return "this is the Kraken!"
}
