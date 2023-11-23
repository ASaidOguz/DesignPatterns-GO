package main

import "fmt"

//this is complex object we need to build
type House struct {
	floors   int
	RoofType string
	rooms    int
}

type HouseBuilder interface {
	SetFloors(int) HouseBuilder
	SetRoofType(string) HouseBuilder
	SetRooms(int) HouseBuilder
	Build() House
}

type ConcreteHouseBuilder struct {
	house House
}

func NewHouseBuilder() HouseBuilder {
	return &ConcreteHouseBuilder{}
}

func (b *ConcreteHouseBuilder) SetFloors(floors int) HouseBuilder {
	b.house.floors = floors
	return b
}

func (b *ConcreteHouseBuilder) SetRoofType(roofType string) HouseBuilder {
	b.house.RoofType = roofType
	return b
}

func (b *ConcreteHouseBuilder) SetRooms(rooms int) HouseBuilder {
	b.house.rooms = rooms
	return b
}

func (b *ConcreteHouseBuilder) Build() House {
	return b.house
}

func main() {
	// Using the Builder to construct a house step by step
	builder := NewHouseBuilder()

	house := builder.SetFloors(2).SetRoofType("Tile").SetRooms(4).Build()

	// Print the constructed house
	fmt.Printf("Constructed House: %+v\n", house)
}

/*
the ConcreteHouseBuilder struct encapsulates the House struct.
This encapsulation helps in the step-by-step construction
 of the House object while keeping track of its attributes within the builder.
*/
