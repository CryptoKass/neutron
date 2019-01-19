package core

import "github.com/veandco/go-sdl2/sdl"

type Viewport sdl.Rect

func (v *Viewport) GetWidth() float32 {
	view := sdl.Rect(*v)
	return float32(view.W)
}

func (v *Viewport) GetHeight() float32 {
	view := sdl.Rect(*v)
	return float32(view.H)
}

func (v *Viewport) Set(w, h float32) {
	view := Viewport(sdl.Rect{0, 0, int32(w), int32(h)})
	v = &view
}
