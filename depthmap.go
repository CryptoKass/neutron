package neutron

import "math/rand"

var depthmap []neutronRef

type neutronRef struct {
	e *Element
}

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
