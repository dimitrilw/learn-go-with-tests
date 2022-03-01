package structmethodinterface

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type RightTriangle struct {
	Base   float64
	Height float64
}

func (rt RightTriangle) Area() float64 {
	return 0.5 * rt.Base * rt.Height
}

func (rt RightTriangle) Perimeter() float64 {
	// A squared + B squared = C squared ... C = sqrt(A squared + B squared)
	c := math.Sqrt(math.Pow(rt.Base, 2) + math.Pow(rt.Height, 2))
	return rt.Base + rt.Height + c
}
