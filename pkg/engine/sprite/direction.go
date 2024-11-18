package sprite

type Direction int

const (
	Up    Direction = 0
	Down            = 1
	Left            = 2
	Right           = 3
)

func isHorizontal(d Direction) bool {
	return d == Left || d == Right
}

func isVertical(d Direction) bool {
	return d == Up || d == Down
}
