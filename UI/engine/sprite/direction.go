package sprite

type Direction int

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

func IsHorizontal(d Direction) bool {
	return d == Left || d == Right
}

func IsVertical(d Direction) bool {
	return d == Up || d == Down
}
