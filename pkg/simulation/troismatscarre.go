package simulation

import (
	"fmt"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type TroisMatsCarre struct {
	sprite.Sprite
}

func NewTroisMatsCarre(position engine.Position) *TroisMatsCarre {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de TroisMatsCarre impossible : %s", err)
		return nil
	}

	return &TroisMatsCarre{*sprite}
}

func (p *TroisMatsCarre) Update() {
	grid.GetInstance().MoveNext(p, p.Direction)
}

func (p *TroisMatsCarre) ToString() string {
	return "this is the TroisMatsCarre!"
}
