package main

import (
	"fmt"
)

func main() {
	var name string = "John"
	var age int = 30
	var isEmployed bool = true

	fmt.Println("Variables in Go")
	fmt.Println("Name:", name, "Type:", fmt.Sprintf("%T", name))
	fmt.Println("Age:", age, "Type:", fmt.Sprintf("%T", age))
	fmt.Println("Employed:", isEmployed, "Type:", fmt.Sprintf("%T", isEmployed))

	var smallValue uint8 = 255
	fmt.Println("Small Value:", smallValue, "Type:", fmt.Sprintf("%T", smallValue))

	var largeValue int64 = 9223372036854775807
	fmt.Println("Large Value:", largeValue, "Type:", fmt.Sprintf("%T", largeValue))

	var temperature float32 = 36.6
	fmt.Println("Temperature:", temperature, "Type:", fmt.Sprintf("%T", temperature))

	var pi float64 = 3.14159
	fmt.Println("Pi:", pi, "Type:", fmt.Sprintf("%T", pi))

	// Default values and some aliases
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("Default Int:", defaultInt, "Type:", fmt.Sprintf("%T", defaultInt))
	fmt.Println("Default Float:", defaultFloat, "Type:", fmt.Sprintf("%T", defaultFloat))
	fmt.Println("Default Bool:", defaultBool, "Type:", fmt.Sprintf("%T", defaultBool))
	fmt.Println("Default String:", defaultString, "Type:", fmt.Sprintf("%T", defaultString))

	// implicit typing
	city := "New York"
	height := 5.9
	isStudent := false

	fmt.Println("City:", city, "Type:", fmt.Sprintf("%T", city))
	fmt.Println("Height:", height, "Type:", fmt.Sprintf("%T", height))
	fmt.Println("Is Student:", isStudent, "Type:", fmt.Sprintf("%T", isStudent))

	// no var style of declaration
	number := 42
	fmt.Println("Number:", number, "Type:", fmt.Sprintf("%T", number))
}
