package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	SdlTexture    *sdl.Texture
	Width, Height float32
}

func (ren *Renderer) CreateTexture(w, h float32) *Texture {
	tex, err := ren.SdlRenderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, int32(w), int32(h))
	if err != nil {
		panic(err)
	}
	return &Texture{
		SdlTexture: tex,
		Width:      w,
		Height:     h,
	}

}

func (ren *Renderer) LoadImage(filename string) *Texture {
	//load iaage to surface
	surf, err := img.Load(filename)
	if err != nil {
		panic(err)
	}
	//load surf to texture
	tex, err := ren.SdlRenderer.CreateTextureFromSurface(surf)
	if err != nil {
		panic(err)
	}
	//get width and height
	_, _, w, h, err := tex.Query()
	if err != nil {
		panic(err)
	}

	return &Texture{
		SdlTexture: tex,
		Width:      float32(w),
		Height:     float32(h),
	}
}

func (ren *Renderer) LoadSdlTexture(tex *sdl.Texture) *Texture {
	_, _, w, h, err := tex.Query()
	if err != nil {
		panic(err)
	}
	return &Texture{
		SdlTexture: tex,
		Width:      float32(w),
		Height:     float32(h),
	}
}
