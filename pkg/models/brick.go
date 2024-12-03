package models

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Brick struct {
	sprite.Sprite
}

func NewBrick(position engine.Position) *Brick {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Brick impossible : %s", err)
		return nil
	}

	return &Brick{*sprite}
}

func (p *Brick) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *Brick) ToString() string {
	return "this is the Brick!"
}
