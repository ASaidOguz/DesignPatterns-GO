package main

import "fmt"

// Strategy interface defines a method for changing algorithms
type Strategy interface {
	ChangeAlgorithm()
}

// Concrete strategy: Zipper
type Zipper struct{}

func (z *Zipper) ChangeAlgorithm() {
	fmt.Println("Zipper algorithm used")
}

// Concrete strategy: WinRARer
type WinRARer struct{}

func (w *WinRARer) ChangeAlgorithm() {
	fmt.Println("WinRARer algorithm used")
}

// Concrete strategy: Gzipper
type Gzipper struct{}

func (g *Gzipper) ChangeAlgorithm() {
	fmt.Println("Gzipper algorithm used")
}

// AlgorithmSelector holds a reference to a strategy
type AlgorithmSelector struct {
	strategy Strategy
}

// Method to change the current strategy
func (as *AlgorithmSelector) SetStrategy(s Strategy) {
	as.strategy = s
}

// Method to perform an action using the current strategy
func (as *AlgorithmSelector) UseStrategy() {
	if as.strategy == nil {
		fmt.Println("Please set a strategy first")
		return
	}
	as.strategy.ChangeAlgorithm()
}

func main() {
	// Initialize AlgorithmSelector
	selector := &AlgorithmSelector{}

	// Set a concrete strategy (e.g., Zipper)
	selector.SetStrategy(&Zipper{})

	// Use the selected strategy
	selector.UseStrategy()

	// Change the strategy to WinRARer and use it
	selector.SetStrategy(&WinRARer{})
	selector.UseStrategy()

	// Change the strategy to Gzipper and use it
	selector.SetStrategy(&Gzipper{})
	selector.UseStrategy()
}

/*
type AlgorithmSelector struct { strategy Strategy } denotes
that AlgorithmSelector is a struct type with a field named strategy
of type Strategy. This field can hold any value that implements the Strategy
interface. It means you can assign any concrete type that satisfies
the Strategy interface to the strategy field within AlgorithmSelector.

--- Check in above for nil needed in case no strategy choosen and my
lead null pointer error...
*/
