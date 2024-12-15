package simulation

type Direction int

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

func isHorizontal(d Direction) bool {
	return d == Left || d == Right
}

func isVertical(d Direction) bool {
	return d == Up || d == Down
}

type ShipState int

const (
	Idle      ShipState = 0
	OnFire    ShipState = 1
	Attacking ShipState = 2
)

type KrakenState int

const (
	Hidden     KrakenState = 0
	OutOfWater KrakenState = 1
)
