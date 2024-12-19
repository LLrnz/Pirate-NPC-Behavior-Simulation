package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// ==================[Kraken View]=====================
// Structure acting as a displayable version of the Kraken object
type ViewKraken struct {
	ViewAgent
}

// Initiate the agent textures
func getKrakenSprites() ([]ebiten.Image, error) {
	im1, err1 := utils.NewImageFromPath("assets/kraken.png")
	im2, err2 := utils.NewImageFromPath("assets/kraken_shadow.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Hidden:     *im2,
		simulation.OutOfWater: *im1,
	}

	return res, nil
}

func NewViewKraken(a *simulation.Kraken) (IViewAgent, error) {
	pos := utils.ToPos(a.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getKrakenSprites()

	if err != nil {
		return &ViewKraken{}, err
	}

	agt := ViewAgent{pos, a.Direction(), sprites, 0, ToInfo}
	return &ViewKraken{agt}, nil
}

// ==================[Pirate View]=====================
type ViewPirate struct {
	ViewAgent
}

func getPirateSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/pirateship.png")
	im1, err1 := utils.NewImageFromPath("assets/pirateship_fire.png")
	im2, err2 := utils.NewImageFromPath("assets/pirateship_attack.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Idle:      *im,
		simulation.OnFire:    *im1,
		simulation.Attacking: *im2,
	}

	return res, err
}

func NewViewPirate(s *simulation.Ship) (IViewAgent, error) {
	pos := utils.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getPirateSprites()

	if err != nil {
		return &ViewPirate{}, err
	}

	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
	return &ViewPirate{agt}, nil
}

// ==================[Army View]=====================
type ViewArmy struct {
	ViewAgent
}

func getArmySprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/armyship.png")
	im1, err1 := utils.NewImageFromPath("assets/armyship_fire.png")
	im2, err2 := utils.NewImageFromPath("assets/armyship_attack.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Idle:      *im,
		simulation.OnFire:    *im1,
		simulation.Attacking: *im2,
	}

	return res, err
}

func NewViewArmy(s *simulation.Ship) (IViewAgent, error) {
	pos := utils.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getArmySprites()

	if err != nil {
		return &ViewArmy{}, err
	}

	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
	return &ViewArmy{agt}, nil
}

// ==================[Privateer View]=====================
type ViewPrivateer struct {
	ViewAgent
}

func getPrivateerSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/privateer.png")
	im1, err1 := utils.NewImageFromPath("assets/privateer_fire.png")
	im2, err2 := utils.NewImageFromPath("assets/privateer_attack.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Idle:      *im,
		simulation.OnFire:    *im1,
		simulation.Attacking: *im2,
	}

	return res, err
}

func NewViewPrivateer(s *simulation.Ship) (IViewAgent, error) {
	pos := utils.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getPrivateerSprites()

	if err != nil {
		return &ViewPrivateer{}, err
	}

	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
	return &ViewPrivateer{agt}, nil
}

// ==================[Merchant View]=====================
type ViewMerchant struct {
	ViewAgent
}

func getMerchantSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/merchantship.png")
	im1, err1 := utils.NewImageFromPath("assets/merchantship_fire.png")
	im2, err2 := utils.NewImageFromPath("assets/merchantship_attack.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Idle:      *im,
		simulation.OnFire:    *im1,
		simulation.Attacking: *im2,
	}

	return res, err
}

func NewViewMerchant(s *simulation.Ship) (IViewAgent, error) {
	pos := utils.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getMerchantSprites()

	if err != nil {
		return &ViewMerchant{}, err
	}

	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
	return &ViewMerchant{agt}, nil
}

// ==================[Rackham View]=====================
type ViewRackham struct {
	ViewAgent
}

func getRackhamSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/zoroship.png")
	im1, err1 := utils.NewImageFromPath("assets/zoroship_fire.png")
	im2, err2 := utils.NewImageFromPath("assets/zoroship_attack.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Idle:      *im,
		simulation.OnFire:    *im1,
		simulation.Attacking: *im2,
	}

	return res, err
}

func NewViewRackham(s *simulation.Ship) (IViewAgent, error) {
	pos := utils.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getRackhamSprites()

	if err != nil {
		return &ViewRackham{}, err
	}

	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
	return &ViewRackham{agt}, nil
}

// ==================[Zoro View]=====================
type ViewZoro struct {
	ViewAgent
}

func getZoroSprites() ([]ebiten.Image, error) {
	im, err := utils.NewImageFromPath("assets/zoroship.png")
	im1, err1 := utils.NewImageFromPath("assets/zoroship_fire.png")
	im2, err2 := utils.NewImageFromPath("assets/zoroship_attack.png")

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := []ebiten.Image{
		simulation.Idle:      *im,
		simulation.OnFire:    *im1,
		simulation.Attacking: *im2,
	}

	return res, err
}

func NewViewZoro(s *simulation.Ship) (IViewAgent, error) {
	pos := utils.ToPos(s.Coordinates())

	ToInfo := func() string {
		return ""
	}

	sprites, err := getZoroSprites()

	if err != nil {
		return &ViewZoro{}, err
	}

	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
	return &ViewZoro{agt}, nil
}
