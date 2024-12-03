package models

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Brigantin struct {
	sprite.Sprite
}

func NewBrigantin(position engine.Position) *Brigantin {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Brigantin impossible : %s", err)
		return nil
	}

	return &Brigantin{*sprite}
}

func (p *Brigantin) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *Brigantin) ToString() string {
	return "this is the Brigantin!"
}
