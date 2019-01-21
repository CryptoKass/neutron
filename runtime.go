package neutron

import (
	"fmt"

	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	Viewport      *core.Viewport
	fps           uint32
	renderer      *core.Renderer
	screenTexture *core.Texture
	window        *sdl.Window

	destroyQue []*neutronRef
)

func (e *engine) tick() {

	for {
		start := sdl.GetTicks()
		if e.pollSdl() {
			return
		}

		//clear the render & set new target
		renderer.SdlRenderer.Clear()
		renderer.SdlRenderer.SetRenderTarget(screenTexture.SdlTexture)

		//Draw to texture
		e.UpdateAll()
		e.DrawAll()

		//copy texture to renderer
		renderer.SdlRenderer.SetRenderTarget(nil)
		renderer.SdlRenderer.Copy(screenTexture.SdlTexture, nil, nil)

		//show the screen
		renderer.SdlRenderer.Present()

		//cleanup destroy que
		Handle(checkCollisions())
		e.DestroyQue()

		frametime := sdl.GetTicks() - start
		if frametime < 1000/fps {
			sdl.Delay((1000 / fps) - frametime)
		}
	}
}

func (e *engine) pollSdl() bool {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			window.Destroy()
			sdl.Quit()
			return true
		}
	}
	return false
}

func (e *engine) UpdateAll() {
	for _, elem := range depthmap {
		if !elem.e.Active {
			continue
		}
		err := elem.e.update()
		if err != nil {
			fmt.Println("error updating element", elem.e.Id, err)
			return
		}
	}
}

func (e *engine) DrawAll() {
	for _, elem := range depthmap {
		if !elem.e.Active {
			continue
		}
		err := elem.e.draw(renderer)
		if err != nil {
			fmt.Println("error drawing element", elem.e.Id, err)
			return
		}
	}
}

func (e *engine) DestroyQue() {
	for _, ref := range destroyQue {
		if ref != nil {
			ref.e.destroy()
		}
	}
	destroyQue = make([]*neutronRef, 0)
}

func (e *engine) GetRenderer() *core.Renderer {
	return renderer
}
