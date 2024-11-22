package vector

import "math"

type Vector [3]float32

func (v Vector) X() float32 {
	return v[0]
}

func (v Vector) Y() float32 {
	return v[1]
}

func (v Vector) Z() float32 {
	return v[2]
}

// Length returns Euclidean norm (length) of the vector.
func (v Vector) Length() float64 {
	return math.Sqrt(
		float64(v[0]*v[0]) + float64(v[1]*v[1]) + float64(v[2]*v[2]),
	)
}

func (v Vector) Normalize() Vector {
	length := v.Length()
	if math.Abs(length) < 1e-6 {
		return [3]float32{0.0, 0.0, 1e-6}
	}

	length = 1.0 / length

	return Vector{
		float32(float64(v[0]) * length),
		float32(float64(v[1]) * length),
		float32(float64(v[2]) * length),
	}
}

func (v Vector) Distance(other Vector) float64 {
	return math.Sqrt(
		float64((v[0]-other[0])*(v[0]-other[0]) +
			(v[1]-other[1])*(v[1]-other[1]) +
			(v[2]-other[2])*(v[2]-other[2])),
	)
}

func (v Vector) IsZero() bool {
	if math.Abs(float64(v[0])) > 1e-6 {
		return false
	}

	if math.Abs(float64(v[1])) > 1e-6 {
		return false
	}

	if math.Abs(float64(v[2])) > 1e-6 {
		return false
	}

	return true
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		v[0] + other[0],
		v[1] + other[1],
		v[2] + other[2],
	}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{
		v[0] - other[0],
		v[1] - other[1],
		v[2] - other[2],
	}
}

func (v Vector) Mul(scalar float32) Vector {
	return Vector{
		v[0] * scalar,
		v[1] * scalar,
		v[2] * scalar,
	}
}

func (v Vector) Div(scalar float32) Vector {
	return Vector{
		v[0] / scalar,
		v[1] / scalar,
		v[2] / scalar,
	}
}

// Dot returns dot product of two vectors.
func (v Vector) Dot(other Vector) float32 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}

func (v Vector) Cross(other Vector) Vector {
	return Vector{
		v[1]*other[2] - v[2]*other[1],
		v[2]*other[0] - v[0]*other[2],
		v[0]*other[1] - v[1]*other[0],
	}
}

func (v Vector) Right() Vector {
	return Vector{v[1], -v[0], 0}
}
