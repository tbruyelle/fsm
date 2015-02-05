package fsm

import (
	"math"
)

type Vector struct {
	X, Y float32
}

func NewVector(fromx, fromy, tox, toy float32) *Vector {
	v := new(Vector)
	v.X = tox - fromx
	v.Y = toy - fromy
	return v
}

func (v *Vector) Length() float32 {
	return float32(math.Sqrt(float64(v.LengthSqr())))
}

func (v *Vector) LengthSqr() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector) Normalize() {
	l := v.Length()
	v.X = v.X / l
	v.Y = v.Y / l
}
