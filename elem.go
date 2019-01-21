package neutron

import (
	"fmt"
	"reflect"

	"github.com/CryptoKass/neutron/core"
	"github.com/CryptoKass/neutron/mathf"
)

// -----------------------------------------------------------|
// REF:Element Struct ----------------------------------------|
// -----------------------------------------------------------|

// Element should be instantiated using NewElement
type Element struct {
	Id         string
	Name       string
	Position   mathf.Vector
	Rotation   float32
	Active     bool
	Components []Component
	colliders  []Collider
	depth      float32
}

// NewElement will create and instantiate a new Element.
func NewElement() *Element {
	elem := Element{
		Id:         GUID(),
		Position:   mathf.Vector{0, 0},
		Rotation:   0,
		Active:     true,
		depth:      0,
		Components: make([]Component, 0),
	}
	//add to elements
	if depthmap == nil {
		//elements = make([]*Element, 0)
		depthmap = make([]neutronRef, 0)
	}
	//elements = append(elements, &elem)
	depthmap = append(depthmap, neutronRef{e: &elem})
	return &elem
}

// FindElement will find an element from the application depth map,
// using the Elements id. If no matching element is found it will
// return nil.
func FindElement(id string) *Element {
	for _, elem := range depthmap {
		if id == elem.e.Id {
			return elem.e
		}
	}
	return nil
}

// -----------------------------------------------------------|
// REF:Element methods ---------------------------------------|
// -----------------------------------------------------------|

// Destroy will remove the element from the scene.
func (elem *Element) destroy() {
	i := elem.GetIndex()
	if i < 0 {
		return
	}
	depthmap = append(depthmap[:i], depthmap[i+1:]...)
}

func (elem *Element) Destroy() {
	destroyQue = append(destroyQue, &neutronRef{elem})
}

// SetDepth set this elements deph. The depth is to order the
// elements during rendering; Higher depth values will appear ontop.
func (elem *Element) SetDepth(z float32) {
	elem.depth = z
	quicksortdepth(depthmap)
}

// GetDepth will return this elements depth, used to order
// elements during rendering; Higher depth values will appear ontop.
func (elem *Element) GetDepth() float32 {
	return elem.depth
}

// AddComponent will add a new Component to this element.
func (elem *Element) AddComponent(new Component) {
	//check comp doesnt already exists
	for _, comp := range elem.Components {
		if reflect.TypeOf(comp) == reflect.TypeOf(new) {
			panic(
				fmt.Sprintf("component duplication error: each element can only have one of each type of component\n"+
					"	 details: attempted to add %t to %s", new, elem.Id),
			)
		}
	}
	elem.Components = append(elem.Components, new)
}

// AddCollider will add a new collider
func (elem *Element) AddCollider(col Collider) {
	col.Circle.Center = elem.Position
	elem.colliders = append(elem.colliders, col)
}

// GetIndex will return the this elements order in the
// depthmap.
func (elem *Element) GetIndex() int {
	for i, e := range depthmap {
		if e.e.Id == elem.Id {
			return i
		}
	}
	return -1
}

// GetComponent Will find a the component matching the type provided.
// The returned component should be cast to a pointer of the
// correct component type.
func (elem *Element) GetComponent(comp Component) Component {
	typ := reflect.TypeOf(comp)
	for _, c := range elem.Components {
		if reflect.TypeOf(c) == typ {
			return c
		}
	}
	panic(fmt.Sprintf("no component with type %v in %s", typ, elem.Id))
}

// -----------------------------------------------------------|
// REF:Element helpers ---------------------------------------|
// -----------------------------------------------------------|

func (elem *Element) draw(rend *core.Renderer) error {
	for _, comp := range elem.Components {
		err := comp.OnDraw(rend)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *Element) update() error {
	for _, col := range elem.colliders {
		col.Circle.Center = elem.Position
	}

	for _, comp := range elem.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *Element) collision(other *Element) error {
	for _, comp := range elem.Components {
		err := comp.OnCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}
