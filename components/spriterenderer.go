package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
	"github.com/CryptoKass/neutron/mathf"
)

type SpriteRenderer struct {
	Element *neutron.Element
	Texture *core.Texture
	Center  mathf.Vector
}

func NewSpriteRenderer(elem *neutron.Element, rend *core.Renderer, filename string) *SpriteRenderer {
	tex := rend.LoadImage(filename)
	return &SpriteRenderer{
		Element: elem,
		Texture: tex,
		Center:  mathf.Vector{X: tex.Width / 2, Y: tex.Height / 2},
	}
}

func (sr *SpriteRenderer) OnDraw(rend *core.Renderer) error {

	ox := sr.Element.Position.X - (sr.Texture.Width / 2)
	oy := sr.Element.Position.Y - (sr.Texture.Height / 2)

	cx := sr.Element.Position.X + sr.Center.X
	cy := sr.Element.Position.Y + sr.Center.Y

	rend.DrawSprite(sr.Texture, ox, oy, cx, cy, sr.Element.Rotation)

	return nil
}
