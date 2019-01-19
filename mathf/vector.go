package mathf

type Vector struct {
	X float32
	Y float32
}

func (v *Vector) Add(vec Vector) Vector {
	return Vector{
		v.X + vec.X,
		v.Y + vec.Y,
	}
}

func (v *Vector) Multi(z float32) Vector {
	return Vector{
		v.X * z,
		v.Y * z,
	}
}

func (v *Vector) Magnitude() float32 {
	return Distance(0, 0, v.X, v.Y)
}

func (v *Vector) Normalize() Vector {
	mag := v.Magnitude()
	if mag <= 0.0001 {
		return Vector{0, 0}
	}
	return v.Multi(1 / mag)
}
