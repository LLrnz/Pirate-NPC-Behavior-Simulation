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
	Pos     utils.Position
	Dir     simulation.Direction
	Texture *ebiten.Image
	State   int
	ToInfo  func() string
}

// Base interface for displayable agent
type IViewAgent interface {
	Draw(screen *ebiten.Image)
	ToString() string
	Position() utils.Position
	SetPosition(utils.Position)
}

type BaseViewFactory struct {
	Textures []ebiten.Image
}

type ViewFactory struct {
	Textures map[string]map[simulation.State]ebiten.Image
}

func NewViewFactory() (*ViewFactory, error) {
	krakenSprites, err := getKrakenSprites()
	if err != nil {
		return nil, err
	}

	pirateSprites, err := getPirateSprites()
	if err != nil {
		return nil, err
	}

	privateerSprites, err := getPrivateerSprites()
	if err != nil {
		return nil, err
	}

	armySprites, err := getArmySprites()
	if err != nil {
		return nil, err
	}

	merchantSprites, err := getMerchantSprites()
	if err != nil {
		return nil, err
	}

	zoroSprites, err := getZoroSprites()
	if err != nil {
		return nil, err
	}

	rackhamSprites, err := getRackhamSprites()
	if err != nil {
		return nil, err
	}

	textures := map[string]map[simulation.State]ebiten.Image{
		"Kraken":    krakenSprites,
		"Pirate":    pirateSprites,
		"Privateer": privateerSprites,
		"Army":      armySprites,
		"Merchant":  merchantSprites,
		"Zoro":      zoroSprites,
		"Rackham":   rackhamSprites,
	}

	return &ViewFactory{textures}, nil
}

// Converts a simulation agent to a displayable view of that agent
// Kinda a factory

func (v *ViewFactory) AgentToView(a simulation.Agent) (IViewAgent, error) {
	pos := utils.ToPos(a.Coordinates())

	switch a.AgtType() {
	case "Kraken":
		kraken, _ := a.(*simulation.Kraken)
		texture := v.Textures[a.AgtType()][0]

		return &ViewAgent{pos, kraken.Direction(), &texture, 0, func() string { return kraken.Id() }}, nil

	case "Ship":
		ship, _ := a.(*simulation.Ship)
		texture := v.Textures[ship.ShipType()][simulation.State(ship.ShipState())]
		return &ViewAgent{pos, ship.Direction(), &texture, int(ship.ShipState()), func() string { return ship.Id() }}, nil

	default:
		return nil, fmt.Errorf("unknown AgentType : %s", a.AgtType())
	}
}

// Converts all agents of a slice to displayable views
func (v *ViewFactory) ConvertAllAgentsToView(agents []simulation.Agent) ([]IViewAgent, error) {
	size := len(agents)
	views := make([]IViewAgent, size)

	for i, agt := range agents {
		view, err := v.AgentToView(agt)

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

	//flip horizontal
	if v.Dir == simulation.Left {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(v.Texture.Bounds().Dx()), 0)
	}

	// //flip vertical
	// if v.Dir == simulation.Down {
	// 	op.GeoM.Scale(1, -1)
	// 	op.GeoM.Translate(0, float64(v.Textures[v.State].Bounds().Dy()))
	// }

	op.GeoM.Translate(float64(v.Pos.X), (float64(v.Pos.Y)))
	screen.DrawImage(v.Texture, op)
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
