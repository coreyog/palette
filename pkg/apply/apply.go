package apply

import (
	"image"

	"github.com/coreyog/palette/pkg/gen"
	"github.com/coreyog/palette/pkg/model"
)

func ApplyPaletteToTemplate(template *image.NRGBA, palette model.Palette) (img *image.NRGBA, err error) {
	img = image.NewNRGBA(template.Bounds())

	ref := gen.ReferencePalette()

	for y := 0; y < template.Bounds().Dy(); y++ {
		for x := 0; x < template.Bounds().Dx(); x++ {
			c := template.At(x, y)
			index := ref.IndexOf(c)

			if index != -1 {
				img.Set(x, y, palette[index])
			}
		}
	}

	return img, nil
}
