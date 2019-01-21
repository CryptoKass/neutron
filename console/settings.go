package console

import (
	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//Settings Go
var (
	Key               = int8(sdl.SCANCODE_F11)
	GameLogOpen       = true
	GameLogLineCount  = 12
	Font, fontLoadErr *ttf.Font
	Color             core.Color
)
