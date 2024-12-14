package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/engine"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// ==================[Kraken View]=====================
type ViewKraken struct {
	ViewAgent
}

func getKrakenSprites() []ebiten.Image {
	return nil
}

func NewViewKraken(a simulation.Agent) IViewAgent {
	pos := engine.ToPos(a.Coordinates())

	ToInfo := func() string {
		return ""
	}

	agt := ViewAgent{pos, getKrakenSprites(), 0, ToInfo}
	return &ViewKraken{agt}
}

// ==================[Pirate View]=====================
type ViewPirate struct {
	ViewAgent
}

func getPirateSprites() []ebiten.Image {
	return nil
}

func NewViewPirate(s *simulation.Ship) IViewAgent {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	agt := ViewAgent{pos, getPirateSprites(), 0, ToInfo}
	return &ViewPirate{agt}
}

// ==================[Army View]=====================
type ViewArmy struct {
	ViewAgent
}

func getArmySprites() []ebiten.Image {
	return nil
}

func NewViewArmy(s *simulation.Ship) IViewAgent {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}
	agt := ViewAgent{pos, getArmySprites(), 0, ToInfo}
	return &ViewArmy{agt}
}

// ==================[Privateer View]=====================
type ViewPrivateer struct {
	ViewAgent
}

func getPrivateerSprites() []ebiten.Image {
	return nil
}

func NewViewPrivateer(s *simulation.Ship) IViewAgent {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}
	agt := ViewAgent{pos, getPrivateerSprites(), 0, ToInfo}
	return &ViewPrivateer{agt}
}

// ==================[Merchant View]=====================
type ViewMerchant struct {
	ViewAgent
}

func getMerchantSprites() []ebiten.Image {
	return nil
}

func NewViewMerchant(s *simulation.Ship) IViewAgent {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}
	agt := ViewAgent{pos, getMerchantSprites(), 0, ToInfo}
	return &ViewMerchant{agt}
}

// ==================[Rackham View]=====================
type ViewRackham struct {
	ViewAgent
}

func getRackhamSprites() []ebiten.Image {
	return nil
}

func NewViewRackham(s *simulation.Ship) IViewAgent {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}
	agt := ViewAgent{pos, getRackhamSprites(), 0, ToInfo}
	return &ViewRackham{agt}
}

// ==================[Zoro View]=====================
type ViewZoro struct {
	ViewAgent
}

func getZoroSprites() []ebiten.Image {
	return nil
}

func NewViewZoro(s *simulation.Ship) IViewAgent {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}
	agt := ViewAgent{pos, getZoroSprites(), 0, ToInfo}
	return &ViewZoro{agt}
}
