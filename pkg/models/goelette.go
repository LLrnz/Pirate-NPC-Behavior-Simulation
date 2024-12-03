package simulation

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Goelette struct {
	sprite.Sprite
}

func NewGoelette(position engine.Position) *Goelette {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Goelette impossible : %s", err)
		return nil
	}

	return &Goelette{*sprite}
}

func (p *Goelette) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *Goelette) ToString() string {
	return "this is the Goelette!"
}
