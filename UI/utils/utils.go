package utils

import (
	"bytes"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

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

func getImageFromFilePath(path string) (*ebiten.Image, error) {

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

func NewImageFromPath(filePath string) (*ebiten.Image, error) {
	im, err := getImageFromFilePath(filePath)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(im), err
}
