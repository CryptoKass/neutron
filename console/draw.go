package console

import "github.com/CryptoKass/neutron/core"

func DrawGameConole(rend *core.Renderer) {

	if !GameLogOpen {
		return
	}

	ww, hh := rend.SdlRenderer.GetLogicalSize()
	top := float32(hh / 2)
	mid := float32(ww / 2)
	rend.SetRGBA(0, 0, 0, 200)
	rend.DrawRect(0, top, float32(ww), top)

	var x, y, cx, cy float32
	var size = Font.LineSkip()
	var spacing = float32(2)
	for i, line := range gameloglines {
		y = top + spacing + float32(i*size) + spacing
		x = 12
		cx = mid - 6
		cy = y + float32(size/2)
		rend.SetRGBA(255, 255, 255, 255)
		rend.DrawText(line, Font, Color, x, y, cx, cy, 0)
	}
}
