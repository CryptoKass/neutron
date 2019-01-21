package core

import (
	"github.com/CryptoKass/neutron/mathf"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (ren *Renderer) DrawRect(x, y, w, h float32) {
	rect := sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
	ren.SdlRenderer.FillRect(&rect)
}

func (ren *Renderer) DrawEllipse(x, y, w, h float32) {
	var points []sdl.Point

	point := sdl.Point{}
	for xx := x; xx < x+w; xx++ {
		for yy := y; yy < y+h; yy++ {
			if mathf.InsideEllipse(xx, yy, x, y, w, h) {
				point.X = int32(xx)
				point.Y = int32(yy)
				points = append(points, point)
			}
		}
	}

	if len(points) > 1 {
		ren.SdlRenderer.DrawPoints(points)
	}
}

func (ren *Renderer) DrawSprite(tex *Texture, x, y, cx, cy, rotation float32) {

	ren.SdlRenderer.CopyEx(
		tex.SdlTexture,
		&sdl.Rect{X: 0, Y: 0, W: int32(tex.Width), H: int32(tex.Height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(tex.Width), H: int32(tex.Height)},
		float64(rotation),
		&sdl.Point{X: int32(cx), Y: int32(cy)},
		sdl.FLIP_NONE,
	)

}

func (ren *Renderer) DrawText(str string, font *ttf.Font, color Color, x, y, cx, cy, rotation float32) {
	surf, _ := font.RenderUTF8Solid(str, *color.GetSDLColor())
	tex, _ := ren.SdlRenderer.CreateTextureFromSurface(surf)
	ren.DrawSprite(ren.LoadSdlTexture(tex), x, y, cx, cy, rotation)
}
