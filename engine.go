package neutron

import (
	"github.com/CryptoKass/neutron/audio"
	"github.com/CryptoKass/neutron/console"
	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type engine struct {
	FirstDraw bool
}

// Engine is a handle for engine specific methods.
// Notably Engine.Init and Engine.Start
var Engine engine

func (e *engine) Init(width, height int32, nfps uint32) {

	fps = nfps
	//start sdl
	Handle(sdl.Init(sdl.INIT_EVERYTHING))
	Handle(ttf.Init())
	//Handle(mix.Init(mix.INIT_OGG))

	//create window and renderer
	_wind, _ := sdl.CreateWindow("something",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height,
		sdl.WINDOW_SHOWN)

	//assign renderer and texture
	window = _wind

	//assign renderer
	_rend, _ := sdl.CreateRenderer(window,
		-1, sdl.RENDERER_ACCELERATED)
	renderer = &core.Renderer{SdlRenderer: _rend}

	_view := core.Viewport(renderer.SdlRenderer.GetViewport())
	Viewport = &_view
	screenTexture = renderer.CreateTexture(Viewport.GetWidth(), Viewport.GetHeight())
	renderer.SetBlend(sdl.BLENDMODE_BLEND)

}

func (e *engine) Start() {
	//start sdl
	defer sdl.Quit()

	//Initialize
	renderer.SdlRenderer.SetRenderTarget(screenTexture.SdlTexture)
	renderer.SetBlend(sdl.BLENDMODE_BLEND)

	//begin console
	console.Init()
	console.Log(" Neutron Game Engine - Copyright Kassius Barker-Monbelly 2019 ")
	console.Log(" ------------------------------------------------------------ ")
	console.Log(" VER: PRE-ALPHA | LICENSE:FREE ")

	//begin audio
	audio.Init()
	startup := audio.LoadAudioClip("./res/startup.wav")
	startup.Play()

	//begin update cycle
	console.Log(fps)
	go e.updateloop()
	e.drawLoop()

}

func (e *engine) SetTitle(title string) {
	window.SetTitle(title)
}
