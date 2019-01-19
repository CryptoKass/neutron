package console

import (
	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/ttf"
)

//Settings Go
var (
	GameLogOpen       = true
	GameLogLineCount  = 12
	Font, fontLoadErr *ttf.Font
	Color             core.Color
)
