package mathf

import (
	"math/rand"
)

func RandRange(min, max float32) float32 {
	return min + (rand.Float32() * (max - min))
}
