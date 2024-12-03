package sprite

import (
	engine "simulationPirate/pkg/engine"
	imagereader "simulationPirate/pkg/engine/ImageReader"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	position  engine.Position
	Direction Direction
	Hitbox    Rectangle
	Zoom      float64
	Texture   *ebiten.Image
}

func NewSprite(position engine.Position, direction Direction, zoom float64, fichier string) (*Sprite, error) {
	texture, err := imagereader.NewImageFromPath(fichier)
	if err != nil {
		return nil, err
	}

	return &Sprite{position, direction, *NewRectangleFromBounds(position.Point, texture.Bounds()), zoom, texture}, nil
}

func (s1 Sprite) Collision(s2 Sprite) bool {
	return ((s1.Hitbox.pointHG.X <= s2.Hitbox.pointHG.X) &&
		(s2.Hitbox.pointHG.X <= s1.Hitbox.pointBD.X) &&

		(s1.Hitbox.pointHG.Y <= s2.Hitbox.pointHG.Y) &&
		(s2.Hitbox.pointHG.Y <= s1.Hitbox.pointBD.Y)) ||

		((s1.Hitbox.pointHG.X <= s2.Hitbox.pointBD.X) &&
			(s2.Hitbox.pointBD.X <= s1.Hitbox.pointBD.X) &&

			(s1.Hitbox.pointHG.Y <= s2.Hitbox.pointBD.Y) &&
			(s2.Hitbox.pointBD.Y <= s1.Hitbox.pointBD.Y))
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//flip horizontal
	if s.Direction == Left {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(s.Texture.Bounds().Dx()), 0)
	}

	//flip vertical
	if s.Direction == Down {
		op.GeoM.Scale(1, -1)
		op.GeoM.Translate(0, float64(s.Texture.Bounds().Dy()))
	}

	op.GeoM.Translate(float64(s.position.X), (float64(s.position.Y)))
	screen.DrawImage(s.Texture, op)
}

func (s *Sprite) Position() engine.Position {
	return s.position
}

func (s *Sprite) SetPosition(pos engine.Position) {
	s.position = pos
}
