package views

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/engine"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

type ViewAgent struct {
	Pos      engine.Position
	Textures []ebiten.Image
	State    int
	ToInfo   func() string
}

type IViewAgent interface {
	Draw(screen *ebiten.Image)
	ToString() string
	Position() engine.Position
	SetPosition(engine.Position)
}

func AgentToView(a simulation.Agent) (IViewAgent, error) {

	switch a.AgtType() {
	case "Kraken":
		return NewViewKraken(a), nil

	case "Ship":
		ship, _ := a.(*simulation.Ship)
		switch ship.ShipType() {
		case "Pirate":
			return NewViewPirate(ship), nil
		case "Army":
			return NewViewArmy(ship), nil
		case "Privateer":
			return NewViewPrivateer(ship), nil
		case "Merchant":
			return NewViewMerchant(ship), nil
		case "Rackham":
			return NewViewRackham(ship), nil
		case "Zoro":
			return NewViewZoro(ship), nil
		default:
			return nil, fmt.Errorf("unknown ShipType : %s", a.AgtType())
		}

	default:
		return nil, fmt.Errorf("unknown AgentType : %s", a.AgtType())
	}
}

func ConvertAllAgentsToView(agents []simulation.Agent) ([]IViewAgent, error) {
	size := len(agents)
	views := make([]IViewAgent, size)

	for i, agt := range agents {
		view, err := AgentToView(agt)

		if err != nil {
			return nil, err
		}

		views[i] = view
	}

	return views, nil
}

func (v *ViewAgent) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//TODO : To be added ?
	// //flip horizontal
	// if s.Direction == Left {
	// 	op.GeoM.Scale(-1, 1)
	// 	op.GeoM.Translate(float64(s.Texture.Bounds().Dx()), 0)
	// }

	// //flip vertical
	// if s.Direction == Down {
	// 	op.GeoM.Scale(1, -1)
	// 	op.GeoM.Translate(0, float64(s.Texture.Bounds().Dy()))
	// }

	op.GeoM.Translate(float64(v.Pos.X), (float64(v.Pos.Y)))
	screen.DrawImage(&v.Textures[v.State], op)
}

func (v *ViewAgent) ToString() string {
	return v.ToInfo()
}

func (v *ViewAgent) Position() engine.Position {
	return v.Pos
}

func (v *ViewAgent) SetPosition(pos engine.Position) {
	v.Pos = pos
}
