package neutron

import (
	"github.com/CryptoKass/neutron/mathf"
)

type Collider struct {
	Circle *mathf.Circle
}

func collides(c1, c2 *mathf.Circle) bool {
	dist := mathf.Sqrt(
		mathf.Sqr(c1.Center.X-c2.Center.X) +
			mathf.Sqr(c1.Center.Y-c2.Center.Y),
	)
	return dist <= c1.Radius+c2.Radius
}

func checkCollisions() error {
	for i := 0; i < len(depthmap)-1; i++ {
		for j := i + 1; j < len(depthmap); j++ {
			//check both elements are active
			if !depthmap[i].e.Active || !depthmap[j].e.Active {
				continue
			}
			//loop throw colliders
			for _, col1 := range depthmap[i].e.colliders {
				for _, col2 := range depthmap[j].e.colliders {
					if collides(col1.Circle, col2.Circle) {
						err := depthmap[i].e.collision(depthmap[j].e)
						if err != nil {
							return err
						}
						err = depthmap[j].e.collision(depthmap[i].e)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func NewCircleCollider(radius float32) Collider {
	c := mathf.Circle{
		Center: mathf.Vector{X: 0, Y: 0},
		Radius: radius,
	}

	return Collider{
		Circle: &c,
	}
}
