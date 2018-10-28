package shapes

import (
	"math"
	"testing"
)

type shapeTests struct {
	name  string
	shape Shape
	want  float64
}

func testShape(t *testing.T, function func(Shape) float64, tests []shapeTests) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.shape); got != tt.want {
				t.Errorf("Shape.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TODO is it possible to link directly to func on interface without this bridge func?
func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

func TestRect_Perimeter(t *testing.T) {
	testShape(
		t, Perimeter,
		[]shapeTests{
			{"square1", Rect{1, 1}, 4},
			{"square2", Rect{2, 2}, 8},
			{"rect12", Rect{1, 2}, 6},
			{"rect21", Rect{2, 1}, 6},
		})
}

func TestTriangular_Perimeter(t *testing.T) {
	testShape(
		t, Perimeter,
		[]shapeTests{
			{"1,1,1", Triangular{1, 1, 1}, 3},
			{"3,4,5", Triangular{3, 4, 5}, 12},
		})
}

func TestCircle_Perimeter(t *testing.T) {
	testShape(
		t, Perimeter,
		[]shapeTests{
			{"circle1", Circle{1}, math.Pi * 2},
			{"circle2", Circle{2}, math.Pi * 4},
		})
}

//TODO is it possible to link directly to func on interface without this bridge func?
func Area(s Shape) float64 {
	return s.Area()
}

func TestRect_Area(t *testing.T) {
	testShape(
		t, Area,
		[]shapeTests{
			{"square1", Rect{1, 1}, 1},
			{"square2", Rect{2, 2}, 4},
			{"rect12", Rect{1, 2}, 2},
			{"rect21", Rect{2, 1}, 2},
		})
}

func TestTriangular_Area(t *testing.T) {
	testShape(
		t, Area,
		[]shapeTests{
			{"1,1,1", Triangular{1, 1, 1}, math.Sqrt(3) / 4},
			{"3,4,5", Triangular{3, 4, 5}, 6},
		})
}

func TestCircle_Area(t *testing.T) {
	testShape(
		t, Area,
		[]shapeTests{
			{"circle1", Circle{1}, math.Pi * 1},
			{"circle2", Circle{2}, math.Pi * 4},
		})
}
