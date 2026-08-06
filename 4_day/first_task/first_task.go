package main

import (
	"fmt"
	"math"
)

var Pi float64 = math.Pi

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	radius := c.Radius
	squareRadius := radius * radius
	return squareRadius * Pi
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func PrintShapeInfo(x Shape) {
	switch v := x.(type) {
	case Circle:
		fmt.Printf("Это круг с радиусом %v. Его площадь: %v", v.Radius, v.Area())
	case Rectangle:
		fmt.Printf("Это прямоугольник со сторонами %v и %v. Его площадь: %v", v.Height, v.Width, v.Area())
	case Triangle:
		fmt.Printf("Это треугольник. Его площадь: %v", v.Area())
	}
}
func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 3, Height: 4},
	}
	sumArea := float64(0)
	for _, v := range shapes {
		sumArea += v.Area()
	}
	PrintShapeInfo(shapes[2])
	fmt.Print("\n")
	fmt.Println(sumArea)
}
