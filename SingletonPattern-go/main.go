package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Initialized!"}
	})

	return instance
}

func main() {
	//create instance if not exist and keep as single created object in entire
	//code.When ever called GetInstance() always return same object...
	instance1 := GetInstance()
	instance2 := GetInstance()

	fmt.Println("Instance-1:", instance1)
	fmt.Println("Instance-1:", instance1)
	fmt.Println("Instance-1 == Instance-2 ? :", instance1 == instance2)
}
