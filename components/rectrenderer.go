package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
)

type RectRenderer struct {
	Element       *neutron.Element
	Color         core.Color
	Width, Height float32
}

func (rect *RectRenderer) OnDraw(rend *core.Renderer) error {
	ww, hh := rect.Width/2, rect.Height/2
	x1 := rect.Element.Position.X - ww
	y1 := rect.Element.Position.Y - hh
	rend.SetColor(rect.Color)
	rend.DrawRect(x1, y1, rect.Width, rect.Height)
	return nil
}

func (rect *RectRenderer) OnUpdate() error {
	return nil
}

func NewRectRenderer(elem *neutron.Element, w, h float32, col core.Color) *RectRenderer {

	return &RectRenderer{
		Element: elem,
		Color:   col,
		Width:   w,
		Height:  h,
	}
}
