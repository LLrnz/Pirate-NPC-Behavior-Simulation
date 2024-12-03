package models

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Port struct {
	sprite.Sprite
}

func NewPort(position engine.Position) *Port {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Port impossible : %s", err)
		return nil
	}

	return &Port{*sprite}
}

func (p *Port) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *Port) ToString() string {
	return "this is the Port!"
}
