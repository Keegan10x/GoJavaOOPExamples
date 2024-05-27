package example

import (
	"math"
)

// Shape interface
type Shape interface {
	CalcPerimeter() float64
	CalcArea() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Circle implements Shape interface
func (c Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Square struct
type Square struct {
	Side float64
}

// Square implements Shape interface
func (s Square) CalcPerimeter() float64 {
	return 4 * s.Side
}

func (s Square) CalcArea() float64 {
	return s.Side * s.Side
}
