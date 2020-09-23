package main

import (
	"testing"
)

func vector(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func TestVector_Add(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"should add correctly", args{vector(1., 1., 1.), vector(2., 2., 2.)}, vector(3., 3., 3.)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add([%v], [%v]) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestVector_Subtract(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"should subtract correctly", args{vector(3., 3., 3.), vector(1., 1., 1.)}, vector(2., 2., 2.)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtract([%v], [%v]) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestVector_MultiplyByScalar(t *testing.T) {
	type args struct {
		a Vector
		f float64
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"should multiply each component by given scalar", args{vector(3., 3., 3.), 2.}, vector(6., 6., 6.)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyByScalar(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("MultiplyByScalar([%v], [%v]) = %v, want %v", tt.args.a, tt.args.f, got, tt.want)
			}
		})
	}
}

func TestVector_Dot(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"dot product of two perpendicular vectors is 0", args{vector(1., 0., 0.), vector(0., 1., 0.)}, 0},
		{"dot product of two parallel vectors is 1", args{vector(1., 0., 0.), vector(1., 0., 0.)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Dot([%v], [%v]) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestVector_Length(t *testing.T) {
	type args struct {
		a Vector
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"calculates the length(magnitude) correctly for Vector{3., 0., 0.}", args{vector(3., 0., 0.)}, 3},
		{"calculates the length(magnitude) correctly for Vector{6., 0., 0.}", args{vector(6., 0., 0.)}, 6},
		{"calculates the length(magnitude) correctly for Vector{6., 2., 0.}", args{vector(6., 2., 0.)}, 6.324555320336759},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Length(tt.args.a); got != tt.want {
				t.Errorf("Length([%v]) = %v, want %v", tt.args.a, got, tt.want)
			}
		})
	}
}

func TestVector_Cross(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"cross product returns a vector perpendicular to the other two", args{vector(1., 0., 0.), vector(0., 1., 0.)}, vector(0., 0., 1.)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cross(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Cross([%v], [%v]) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestVector_UnitVector(t *testing.T) {
	type args struct {
		a Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"returns a unit vector from the given vector", args{vector(10., 0., 0.)}, vector(1., 0., 0.)},
		{"returns a unit vector from the given vector", args{vector(0., 10., 0.)}, vector(0., 1., 0.)},
		{"returns a unit vector from the given vector", args{vector(0., 0., 10.)}, vector(0., 0., 1.)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitVector(tt.args.a); got != tt.want {
				t.Errorf("UnitVector([%v]) = %v, want %v", tt.args.a, got, tt.want)
			}
		})
	}
}
