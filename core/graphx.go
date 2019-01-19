package core

import "github.com/veandco/go-sdl2/sdl"

type Renderer struct {
	SdlRenderer *sdl.Renderer
}

func (ren *Renderer) SetColor(col Color) {

	ren.SdlRenderer.SetDrawColor(col.R, col.G, col.B, col.A)
}

func (ren *Renderer) SetRGBA(r, g, b, a uint8) {
	ren.SdlRenderer.SetDrawColor(r, g, b, a)
}

func (ren *Renderer) SetBlend(bm sdl.BlendMode) {
	ren.SdlRenderer.SetDrawBlendMode(bm)
}
