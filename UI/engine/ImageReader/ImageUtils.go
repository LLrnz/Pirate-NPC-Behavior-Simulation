package imagereader

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func NewImageFromPath(filePath string) (*ebiten.Image, error) {
	im, err := getImageFromFilePath(filePath)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(im), err
}
