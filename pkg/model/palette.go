package model

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

var (
	ErrInvalidPaletteSize = fmt.Errorf("palette image must contain 256 or fewer pixels")
)

type Palette []color.NRGBA

func NewPalette() Palette {
	arr := make([]color.NRGBA, 0, 256)
	arr = append(arr, color.NRGBA{0, 0, 0, 0})
	return Palette(arr)
}

func (p Palette) IndexOf(c color.Color) int {
	rgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	for i, v := range p {
		if v == rgba {
			return i
		}
	}

	return -1
}

func (p Palette) Add(c color.Color) (bool, Palette) {
	rgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	if len(p) < 256 && p.IndexOf(rgba) == -1 {
		p = append(p, rgba)
		return true, p
	}

	return false, p
}

func (p Palette) Remove(c color.Color) Palette {
	rgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	for i, v := range p {
		if v == rgba {
			p = append(p[:i], p[i+1:]...)
			break
		}
	}

	return p
}

func (p Palette) ToImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, 16, 16))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.NRGBA{0, 0, 0, 255}}, image.Point{}, draw.Src)

	for i, c := range p {
		y := i / 16
		x := i % 16

		img.Set(x, y, c)
	}

	return img
}

func LoadPalette(img *image.NRGBA) (Palette, error) {
	p := NewPalette()

	if img.Bounds().Dx()*img.Bounds().Dy() > 256 {
		return p, ErrInvalidPaletteSize
	}

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			c := img.At(x, y)
			_, p = p.Add(c)
		}
	}

	return p, nil
}
