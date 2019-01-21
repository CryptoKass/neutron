package neutron

import "github.com/CryptoKass/neutron/core"

type Component interface {
	OnUpdate() error
	OnDraw(rend *core.Renderer) error
	OnCollision(other *Element) error
}
