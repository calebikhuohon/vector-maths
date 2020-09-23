package main

import "math"

type Vector struct {
	X, Y, Z float64
}

func Add(a, b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Subtract(a, b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func MultiplyByScalar(a Vector, f float64) Vector {
	return Vector{
		X: a.X * f,
		Y: a.Y * f,
		Z: a.Z * f,
	}
}

func Dot(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Length(a Vector) float64 {
	return math.Sqrt(Dot(a, a))
}

func Cross(a, b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func UnitVector(a Vector) Vector {
	return MultiplyByScalar(a, 1./Length(a))
}
