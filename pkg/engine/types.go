package engine

import "github.com/hajimehoshi/ebiten/v2"

type Updatable interface {
	Update()
}

func DoUpdate(u Updatable) {
	u.Update()
}

type Drawable interface {
	Draw()
}

type Object interface {
	Draw(*ebiten.Image)
	Update()
}

// type Object struct {
// 	position image.Point
// 	Update   func()
// }

// type TexturedObject struct {
// 	Object
// 	sprite sprite.Sprite
// 	//GetSprite func()
// }

// func NewObject(position image.Point, Update func()) *Object {
// 	return &Object{position, Update}
// }

// func NewTexturesdObject(position image.Point, sprite sprite.Sprite, Update func(), Draw func()) *TexturedObject {
// 	return &TexturedObject{*NewObject(position, Update), sprite, Draw}
// }
