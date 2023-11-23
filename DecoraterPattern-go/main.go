package main

import "fmt"

// Shape defines the interface for drawing shapes
type Shape interface {
	Draw() string
}

// Component represents the base shape
type Component struct{}

func (c *Component) Draw() string {
	return "Drawing a basic shape"
}

// Decorator represents the decorator struct
type Decorator struct {
	Shape // Embedded Shape interface
	extra string
}

func (d *Decorator) Draw() string {
	return d.Shape.Draw() + ", " + d.extra
}

func main() {
	// Create a base component
	baseComponent := &Component{}

	// Wrap the base component with decorators
	coloredComponent := &Decorator{Shape: baseComponent, extra: "colored"}
	largeColoredComponent := &Decorator{Shape: coloredComponent, extra: "large"}

	// Display the output of drawing each component
	fmt.Println(baseComponent.Draw())         // Drawing a basic shape
	fmt.Println(coloredComponent.Draw())      // Drawing a basic shape, colored
	fmt.Println(largeColoredComponent.Draw()) // Drawing a basic shape, colored, large
}
