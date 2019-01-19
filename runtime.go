package neutron

import (
	"fmt"

	"github.com/CryptoKass/neutron/console"
	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	Viewport      *core.Viewport
	fps           uint32
	renderer      *core.Renderer
	screenTexture *core.Texture
	window        *sdl.Window
)

func (e *engine) handleInput() {
	var running = true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				window.Destroy()
				break
			}
		}
	}
}

func (e *engine) tick() {
	for {
		start := sdl.GetTicks()

		//clear the render & set new target
		renderer.SdlRenderer.Clear()
		renderer.SdlRenderer.SetRenderTarget(screenTexture.SdlTexture)
		//Draw to texture
		e.DrawAll()

		//copy texture to renderer
		renderer.SdlRenderer.SetRenderTarget(nil)
		renderer.SdlRenderer.Copy(screenTexture.SdlTexture, nil, nil)

		//show the screen
		renderer.SdlRenderer.Present()

		frametime := sdl.GetTicks() - start
		if frametime < 1000/fps {
			sdl.Delay((1000 / fps) - frametime)
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
	console.DrawGameConole(renderer)
}

func (e *engine) GetRenderer() *core.Renderer {
	return renderer
}
