package neutron

import (
	"testing"

	"github.com/CryptoKass/neutron/mathf"
)

func TestCollision(t *testing.T) {
	c1 := mathf.Circle{Center: mathf.Vector{X: 0, Y: 0}, Radius: 5}
	c2 := mathf.Circle{Center: mathf.Vector{X: 1000, Y: 1000}, Radius: 5}

	if collides(&c1, &c2) {
		t.Error("wtf collision")
	}

}
