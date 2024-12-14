package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/engine"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// ==================[Kraken View]=====================
type ViewKraken struct {
	ViewAgent
}

func getKrakenSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/kraken.png")
	return []ebiten.Image{*im}, err
}

func NewViewKraken(a simulation.Agent) (IViewAgent, error) {
	pos := engine.ToPos(a.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getKrakenSprites()

	if err != nil {
		return &ViewKraken{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewKraken{agt}, nil
}

// ==================[Pirate View]=====================
type ViewPirate struct {
	ViewAgent
}

func getPirateSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/pirateship.png")
	return []ebiten.Image{*im}, err
}

func NewViewPirate(s *simulation.Ship) (IViewAgent, error) {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getPirateSprites()

	if err != nil {
		return &ViewPirate{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewPirate{agt}, nil
}

// ==================[Army View]=====================
type ViewArmy struct {
	ViewAgent
}

func getArmySprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/armyship.png")
	return []ebiten.Image{*im}, err
}

func NewViewArmy(s *simulation.Ship) (IViewAgent, error) {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getArmySprites()

	if err != nil {
		return &ViewArmy{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewArmy{agt}, nil
}

// ==================[Privateer View]=====================
type ViewPrivateer struct {
	ViewAgent
}

func getPrivateerSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/huntership.png")
	return []ebiten.Image{*im}, err
}

func NewViewPrivateer(s *simulation.Ship) (IViewAgent, error) {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getPrivateerSprites()

	if err != nil {
		return &ViewPrivateer{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewPrivateer{agt}, nil
}

// ==================[Merchant View]=====================
type ViewMerchant struct {
	ViewAgent
}

func getMerchantSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/merchantship.png")
	return []ebiten.Image{*im}, err
}

func NewViewMerchant(s *simulation.Ship) (IViewAgent, error) {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getMerchantSprites()

	if err != nil {
		return &ViewMerchant{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewMerchant{agt}, nil
}

// ==================[Rackham View]=====================
type ViewRackham struct {
	ViewAgent
}

func getRackhamSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/zoroship.png")
	return []ebiten.Image{*im}, err
}

func NewViewRackham(s *simulation.Ship) (IViewAgent, error) {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getRackhamSprites()

	if err != nil {
		return &ViewRackham{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewRackham{agt}, nil
}

// ==================[Zoro View]=====================
type ViewZoro struct {
	ViewAgent
}

func getZoroSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/zoroship.png")
	return []ebiten.Image{*im}, err
}

func NewViewZoro(s *simulation.Ship) (IViewAgent, error) {
	pos := engine.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getZoroSprites()

	if err != nil {
		return &ViewZoro{}, err
	}

	agt := ViewAgent{pos, sprites, 0, ToInfo}
	return &ViewZoro{agt}, nil
}
