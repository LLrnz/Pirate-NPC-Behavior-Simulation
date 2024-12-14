package views

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// Base structure for displayable agent
// it only stores its position, textures, state
// and a way to communicate its info
type ViewAgent struct {
	Pos      utils.Position
	Textures []ebiten.Image
	State    int
	ToInfo   func() string
}

// Base interface for displayable agent
type IViewAgent interface {
	Draw(screen *ebiten.Image)
	ToString() string
	Position() utils.Position
	SetPosition(utils.Position)
}

// Converts a simulation agent to a displayable view of that agent
// Kinda a factory
func AgentToView(a simulation.Agent) (IViewAgent, error) {

	switch a.AgtType() {
	case "Kraken":
		return NewViewKraken(a)

	case "Ship":
		ship, _ := a.(*simulation.Ship)
		switch ship.ShipType() {
		case "Pirate":
			return NewViewPirate(ship)
		case "Army":
			return NewViewArmy(ship)
		case "Privateer":
			return NewViewPrivateer(ship)
		case "Merchant":
			return NewViewMerchant(ship)
		case "Rackham":
			return NewViewRackham(ship)
		case "Zoro":
			return NewViewZoro(ship)
		default:
			return nil, fmt.Errorf("unknown ShipType : %s", a.AgtType())
		}

	default:
		return nil, fmt.Errorf("unknown AgentType : %s", a.AgtType())
	}
}

// Converts all agents of a slice to displayable views
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

// Draws a displayable agent on screen
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

// Will be used by the info panel to display the agent's info
func (v *ViewAgent) ToString() string {
	return v.ToInfo()
}

// Returns the postition on screen of the displayable agent
func (v *ViewAgent) Position() utils.Position {
	return v.Pos
}

// Sets the position on screen of the displayable agent
func (v *ViewAgent) SetPosition(pos utils.Position) {
	v.Pos = pos
}
