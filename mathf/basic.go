package mathf

import "math"

func Distance(x1, y1, x2, y2 float32) float32 {
	return float32(
		math.Sqrt(
			math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2),
		),
	)
}

func Sqr(z float32) float32 {
	return float32(
		math.Pow(float64(z), 2),
	)
}

func Sign(z float32) float32 {
	if z < 0 {
		return -1
	}
	return 1
}

func Abs(z float32) float32 {
	return float32(
		math.Abs(float64(z)),
	)
}

func Clamp(z, min, max float32) float32 {
	if z < min {
		return min
	}
	if z > max {
		return max
	}
	return z
}
