package utils

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"os"

	emage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

// Loads a font face used by UI elements
func LoadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}

// Loads an ebiten.Image from image file to create agents sprites
func NewImageFromPath(path string) (*ebiten.Image, error) {

	//Open the image
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	//Decode image file
	image, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	//Make the image be an ebiten.Image
	imageEbiten := ebiten.NewImageFromImage(image)

	return imageEbiten, err
}

// Loads nineslices as a button texture
func LoadButtonImage() (*widget.ButtonImage, error) {
	idle := emage.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := emage.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := emage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
