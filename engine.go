package neutron

import (
	"fmt"

	"github.com/CryptoKass/neutron/audio"
	"github.com/CryptoKass/neutron/console"
	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type engine struct{}

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
	_wind, _rend, _ := sdl.CreateWindowAndRenderer(
		width,
		height,
		sdl.WINDOW_SHOWN)

	//assign renderer and texture
	window = _wind
	renderer = &core.Renderer{SdlRenderer: _rend}
	_view := core.Viewport(renderer.SdlRenderer.GetViewport())
	fmt.Println(_view)
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

	//todo elements init

	//begin update cycle
	go e.tick()

	//Handle Input
	e.handleInput()

}
