package neutron

import "github.com/CryptoKass/neutron/core"

type Component interface {
	//onUpdate() error
	OnDraw(rend *core.Renderer) error
}
