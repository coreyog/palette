package util

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func ReadImage(path string) (img *image.NRGBA, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	unenc, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return Convert(unenc), nil
}

func WriteImage(path string, img *image.NRGBA) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func Convert(img image.Image) *image.NRGBA {
	nrgba := image.NewNRGBA(img.Bounds())
	draw.Draw(nrgba, nrgba.Bounds(), img, img.Bounds().Min, draw.Src)

	return nrgba
}
