package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (col *Color) GetSDLColor() *sdl.Color {
	return &sdl.Color{
		R: col.R,
		G: col.G,
		B: col.B,
		A: col.A,
	}
}
