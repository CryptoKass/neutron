package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
	"github.com/CryptoKass/neutron/input"
	"github.com/CryptoKass/neutron/mathf"
	"github.com/veandco/go-sdl2/sdl"
)

type Controller struct {
	Element      *neutron.Element
	Velocity     mathf.Vector
	Speed        float32
	Acceleration float32
	Decay        float32
}

func NewController(elem *neutron.Element, speed, accel, decay float32) *Controller {
	return &Controller{
		Element:      elem,
		Speed:        speed,
		Decay:        decay,
		Acceleration: accel,
	}
}

func (c *Controller) OnUpdate() error {

	//get input X
	if input.IsKeyDown(sdl.SCANCODE_A) {
		c.Velocity.X -= c.Acceleration
	} else if input.IsKeyDown(sdl.SCANCODE_D) {
		c.Velocity.X += c.Acceleration
	} else {
		c.Velocity.X *= c.Decay
	}
	c.Velocity.X = mathf.Clamp(c.Velocity.X, -1, 1)

	//get input Y
	if input.IsKeyDown(sdl.SCANCODE_W) {
		c.Velocity.Y -= c.Acceleration
	} else if input.IsKeyDown(sdl.SCANCODE_S) {
		c.Velocity.Y += c.Acceleration
	} else {
		c.Velocity.Y *= c.Decay
	}
	c.Velocity.Y = mathf.Clamp(c.Velocity.Y, -1, 1)

	//normalize vector
	//norm := c.Velocity.Normalize() c.Velocity.Magnitude()

	c.Element.Position.X += c.Velocity.X * c.Speed
	c.Element.Position.Y += c.Velocity.Y * c.Speed

	//log.Printf("rect at [%.1f,%.1f] moving at [%.1f,%.1f] %.1f", c.Element.Position.X, c.Element.Position.Y,
	//	norm.X, norm.Y, norm.Magnitude())
	return nil
}

func (c *Controller) OnDraw(rend *core.Renderer) error {
	return nil
}
