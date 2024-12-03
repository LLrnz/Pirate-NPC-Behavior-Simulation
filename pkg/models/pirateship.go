package models

import (
	"fmt"
	"testGUI/pkg/engine"
	"testGUI/pkg/engine/sprite"
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

func (p *Pirateship) ToString() string {
	return "this is the pirateship!"
}

func (p *Pirateship) Update() {

}
