package main

import "fmt"

// Vehicle is the interface for creating vehicles
type Vehicle interface {
	Drive()
}

type Car struct{}

type Bike struct{}

type Truck struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a Car")
}

func (c *Bike) Drive() {
	fmt.Println("Driving a Bike")
}

func (c *Truck) Drive() {
	fmt.Println("Driving a Truck")
}

type VehicleFactory interface {
	newVehicle() Vehicle
}

type CarFactory struct{}

type BikeFactory struct{}

type TruckFactory struct{}

func (c *CarFactory) newVehicle() Vehicle {
	return &Car{}
}

func (c *BikeFactory) newVehicle() Vehicle {
	return &Bike{}
}

func (c *TruckFactory) newVehicle() Vehicle {
	return &Truck{}
}
func main() {
	//establish car factory
	carFactory := &CarFactory{}

	//establish bike factory
	bikeFactory := &BikeFactory{}

	//establish truck factory
	truckFactory := &TruckFactory{}

	//manufacture car
	car := carFactory.newVehicle()

	//manufacture bike
	bike := bikeFactory.newVehicle()
	//manufacture truck
	truck := truckFactory.newVehicle()

	//let's drive our car
	car.Drive()

	//let's drive our bike
	bike.Drive()

	//let's drive our truck
	truck.Drive()
}
