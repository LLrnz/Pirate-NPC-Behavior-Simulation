package sprite

import (
	"image"
	imagereader "simulationPirate/pkg/ImageReader"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	position  *image.Point
	direction Direction
	hitbox    Rectangle
	zoom      float64
	texture   *ebiten.Image
}

func NewSprite(position *image.Point, direction Direction, zoom float64, fichier string) (*Sprite, error) {
	texture, err := imagereader.NewImageFromPath(fichier)
	if err != nil {
		return nil, err
	}
	return &Sprite{position, direction, *NewRectangleFromBounds(*position, texture.Bounds()), zoom, texture}, nil
}

func (s Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//flip horizontal
	if s.direction == Left {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(s.texture.Bounds().Dx()), 0)
	}

	//flip vertical
	if s.direction == Down {
		op.GeoM.Scale(1, -1)
		op.GeoM.Translate(0, float64(s.texture.Bounds().Dy()))
	}

	op.GeoM.Translate(float64(s.position.X), (float64(s.position.Y)))
	screen.DrawImage(s.texture, op)
}

func (s Sprite) Position() image.Point {
	return *s.position
}

func (s Sprite) SetPosition(pos image.Point) {
	s.hitbox.pointBD = pos
}

func (s1 Sprite) Collision(s2 Sprite) bool {
	return ((s1.hitbox.pointHG.X <= s2.hitbox.pointHG.X) &&
		(s2.hitbox.pointHG.X <= s1.hitbox.pointBD.X) &&

		(s1.hitbox.pointHG.Y <= s2.hitbox.pointHG.Y) &&
		(s2.hitbox.pointHG.Y <= s1.hitbox.pointBD.Y)) ||

		((s1.hitbox.pointHG.X <= s2.hitbox.pointBD.X) &&
			(s2.hitbox.pointBD.X <= s1.hitbox.pointBD.X) &&

			(s1.hitbox.pointHG.Y <= s2.hitbox.pointBD.Y) &&
			(s2.hitbox.pointBD.Y <= s1.hitbox.pointBD.Y))
}
