package neutron

import "math/rand"

var depthmap []neutronRef

// This is an internal only reference to element, allowing
// reference copies to be made easily.
type neutronRef struct {
	e *Element
}

// quicksortdepth will apply to quick sort algorithm to the
// depthmap. Ordering the elements based on Element.depth.
func quicksortdepth(a []neutronRef) []neutronRef {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].e.depth < a[right].e.depth {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksortdepth(a[:left])
	quicksortdepth(a[left+1:])

	return a
}

// Instantiate adds an element to the game, unless the element
// already exists int the depth map.
func Instantiate(elem *Element) {
	for _, ref := range depthmap {
		if ref.e.Id != elem.Id {
			return
		}
	}

	//add to elements
	if depthmap == nil {
		//elements = make([]*Element, 0)
		depthmap = make([]neutronRef, 0)
	}

	//elements = append(elements, &elem)
	depthmap = append(depthmap, neutronRef{e: elem})
	depthmap = quicksortdepth(depthmap)
}
