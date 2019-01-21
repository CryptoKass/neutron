package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
)

type EllipseRenderer struct {
	Element       *neutron.Element
	Color         core.Color
	Width, Height float32
}

func (ell *EllipseRenderer) OnDraw(rend *core.Renderer) error {
	ww, hh := ell.Width/2, ell.Height/2
	x1 := ell.Element.Position.X - ww
	y1 := ell.Element.Position.Y - hh
	rend.SetColor(ell.Color)
	rend.DrawEllipse(x1, y1, ell.Width, ell.Height)
	return nil
}

func (ell *EllipseRenderer) OnUpdate() error {
	return nil
}

func (ell *EllipseRenderer) OnCollision(other *neutron.Element) error {
	return nil
}

func NewEllipseRenderer(elem *neutron.Element, w, h float32, col core.Color) *EllipseRenderer {
	return &EllipseRenderer{
		Element: elem,
		Color:   col,
		Width:   w,
		Height:  h,
	}
}
