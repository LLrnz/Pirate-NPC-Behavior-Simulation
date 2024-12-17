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
func getKrakenSprites() (map[simulation.State]ebiten.Image, error) {
	im1, err1 := utils.NewImageFromPath(utils.KrakenIdle)
	im2, err2 := utils.NewImageFromPath(utils.KrakenHidden)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.KrakenHidden:     *im2,
		simulation.KrakenOutOfWater: *im1,
	}

	return res, nil
}

// func NewViewKraken(a *simulation.Kraken, textures map[simulation.State]ebiten.Image) (IViewAgent, error) {
// 	pos := utils.ToPos(a.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, a.Direction(), textures[0], 0, ToInfo}
// 	return &ViewKraken{agt}, nil
// }

// ==================[Pirate View]=====================
type ViewPirate struct {
	ViewAgent
}

func getPirateSprites() (map[simulation.State]ebiten.Image, error) {
	im, err := utils.NewImageFromPath(utils.PirateIdle)
	im1, err1 := utils.NewImageFromPath(utils.PirateFire)
	im2, err2 := utils.NewImageFromPath(utils.PirateAttack)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.ShipIdle:      *im,
		simulation.ShipOnFire:    *im1,
		simulation.ShipAttacking: *im2,
	}

	return res, err
}

// func NewViewPirate(s *simulation.Ship, t map[simulation.State]ebiten.Image) (IViewAgent, error) {
// 	pos := utils.ToPos(s.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
// 	return &ViewPirate{agt}, nil
// }

// ==================[Army View]=====================
type ViewArmy struct {
	ViewAgent
}

func getArmySprites() (map[simulation.State]ebiten.Image, error) {
	im, err := utils.NewImageFromPath(utils.ArmyIdle)
	im1, err1 := utils.NewImageFromPath(utils.ArmyFire)
	im2, err2 := utils.NewImageFromPath(utils.ArmyAttack)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.ShipIdle:      *im,
		simulation.ShipOnFire:    *im1,
		simulation.ShipAttacking: *im2,
	}

	return res, err
}

// func NewViewArmy(s *simulation.Ship) (IViewAgent, error) {
// 	pos := utils.ToPos(s.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
// 	return &ViewArmy{agt}, nil
// }

// ==================[Privateer View]=====================
type ViewPrivateer struct {
	ViewAgent
}

func getPrivateerSprites() (map[simulation.State]ebiten.Image, error) {
	im, err := utils.NewImageFromPath(utils.PrivateerIdle)
	im1, err1 := utils.NewImageFromPath(utils.PrivateerFire)
	im2, err2 := utils.NewImageFromPath(utils.PrivateerAttack)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.ShipIdle:      *im,
		simulation.ShipOnFire:    *im1,
		simulation.ShipAttacking: *im2,
	}

	return res, err
}

// func NewViewPrivateer(s *simulation.Ship) (IViewAgent, error) {
// 	pos := utils.ToPos(s.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
// 	return &ViewPrivateer{agt}, nil
// }

// ==================[Merchant View]=====================
type ViewMerchant struct {
	ViewAgent
}

func getMerchantSprites() (map[simulation.State]ebiten.Image, error) {
	im, err := utils.NewImageFromPath(utils.MerchantIdle)
	im1, err1 := utils.NewImageFromPath(utils.MerchantFire)
	im2, err2 := utils.NewImageFromPath(utils.MerchantAttack)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.ShipIdle:      *im,
		simulation.ShipOnFire:    *im1,
		simulation.ShipAttacking: *im2,
	}

	return res, err
}

// func NewViewMerchant(s *simulation.Ship) (IViewAgent, error) {
// 	pos := utils.ToPos(s.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
// 	return &ViewMerchant{agt}, nil
// }

// ==================[Rackham View]=====================
type ViewRackham struct {
	ViewAgent
}

func getRackhamSprites() (map[simulation.State]ebiten.Image, error) {
	im, err := utils.NewImageFromPath(utils.ZoroIdle)
	im1, err1 := utils.NewImageFromPath(utils.ZoroFire)
	im2, err2 := utils.NewImageFromPath(utils.ZoroAttack)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.ShipIdle:      *im,
		simulation.ShipOnFire:    *im1,
		simulation.ShipAttacking: *im2,
	}

	return res, err
}

// func NewViewRackham(s *simulation.Ship) (IViewAgent, error) {
// 	pos := utils.ToPos(s.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
// 	return &ViewRackham{agt}, nil
// }

// ==================[Zoro View]=====================
type ViewZoro struct {
	ViewAgent
}

func getZoroSprites() (map[simulation.State]ebiten.Image, error) {
	im, err := utils.NewImageFromPath(utils.ZoroIdle)
	im1, err1 := utils.NewImageFromPath(utils.ZoroFire)
	im2, err2 := utils.NewImageFromPath(utils.ZoroAttack)

	if err1 != nil {
		return nil, err1
	} else if err2 != nil {
		return nil, err2
	}

	res := map[simulation.State]ebiten.Image{
		simulation.ShipIdle:      *im,
		simulation.ShipOnFire:    *im1,
		simulation.ShipAttacking: *im2,
	}

	return res, err
}

// func NewViewZoro(s *simulation.Ship) (IViewAgent, error) {
// 	pos := utils.ToPos(s.Coordinates())

// 	ToInfo := func() string {
// 		return ""
// 	}

// 	agt := ViewAgent{pos, s.Direction(), sprites, int(s.ShipState()), ToInfo}
// 	return &ViewZoro{agt}, nil
// }
