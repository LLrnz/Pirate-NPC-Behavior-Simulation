package view

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// getImageFromFilePath returns the eibiten image object from a path
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
