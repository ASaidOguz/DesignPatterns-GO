package main

import "fmt"

//square interface
type Square interface {
	GetWidth() float64
}

// Rectangle interface
type Rectangle interface {
	GetWidth() float64
	GetHeight() float64
}

// SquareImpl struct implementing Square interface
type SquareImpl struct {
	Width float64
}

func (s *SquareImpl) GetWidth() float64 {
	return s.Width
}

// SquareAdapter adapting Square to Rectangle interface
type SquareAdapter struct {
	Square
}

func (sa *SquareAdapter) GetHeight() float64 {
	// For a square, height = width
	return sa.GetWidth()
}

func main() {

	square := SquareImpl{
		Width: 5.00,
	}
	// Adapt the square to a rectangle
	adapter := SquareAdapter{Square: &square}

	// Now, we can use the adapted interface
	fmt.Println("Square Width:", square.GetWidth())
	fmt.Println("Rectangle Width:", adapter.GetWidth())
	fmt.Println("Rectangle Height:", adapter.GetHeight())
}

//Difference between creating instance with pointer and no pointer;
/*
Methods with pointer receivers (like func (c *Circle) area() float32)
modify the struct itself. So, in order to pass the Circle struct
and have it modified, I used a pointer (&Circle{...})
to create an instance of Circle.

When you use a pointer receiver in a method, modifying the struct's state will reflect
those changes outside the method. If you want modifications made to a struct within a function
to be reflected outside that function's scope, you should pass a pointer to that struct.

Struct can be composed of different interfaces so it can connect other structs---
*/
