package engine

import(
    "simulationPirate/pkg/sprite"
	"image"
)


type Object interface{
	position image.Point
	Update()
}


type TexturedObject struct{
	Object
	sprite sprite.Sprite
	Draw()
}