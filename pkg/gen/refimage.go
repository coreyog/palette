package gen

import (
	"image"
	"image/color"

	"github.com/coreyog/palette/pkg/model"
)

func ReferenceImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, 16, 16))

	// step := float64(255) / 16
	// halfStep := step / 2
	ref := ReferencePalette()

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			img.Set(x, y, ref[y*16+x])
		}
	}

	return img
}

func ReferencePalette() model.Palette {
	// don't use constructor, don't want transparent color
	p := model.Palette(make([]color.NRGBA, 0, 256))

	step := float64(256) / 16.0

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			_, p = p.Add(color.NRGBA{
				R: uint8(float64(x) * step),
				G: 0,
				B: uint8(float64(y) * step),
				A: 255,
			})
		}
	}

	return p
}
