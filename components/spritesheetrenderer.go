package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
	"github.com/CryptoKass/neutron/mathf"
)

type SpriteSheetRenderer struct {
	Element  *neutron.Element
	Texture  *core.Texture
	Center   mathf.Vector
	Filename string
}
