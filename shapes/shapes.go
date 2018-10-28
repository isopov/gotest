package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	A, B float64
}

func (r Rect) Area() float64 {
	return r.A * r.B
}

func (r Rect) Perimeter() float64 {
	return r.A*2 + r.B*2
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.R * 2
}

type Triangular struct {
	A, B, C float64
}

func (t Triangular) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t Triangular) Perimeter() float64 {
	return t.A + t.B + t.C
}
