package structmethodinterface

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{10}, hasPerimeter: 62.83185307179586},
		{name: "RightTriangle", shape: RightTriangle{12, 6}, hasPerimeter: 31.41640786499874},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "RightTriangle", shape: RightTriangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
