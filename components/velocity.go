package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
	"github.com/CryptoKass/neutron/mathf"
)

type ConstantVelocity struct {
	Element  *neutron.Element
	Velocity mathf.Vector
}

func (cv *ConstantVelocity) OnDraw(rend *core.Renderer) error {

	return nil
}

func (cv *ConstantVelocity) OnUpdate() error {
	cv.Element.Position = cv.Element.Position.Add(cv.Velocity)
	return nil
}

func NewConstantVelocity(elem *neutron.Element, vec mathf.Vector) *ConstantVelocity {
	return &ConstantVelocity{
		Element:  elem,
		Velocity: vec,
	}
}

func (cv *ConstantVelocity) OnCollision(other *neutron.Element) error {
	return nil
}
