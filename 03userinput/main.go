package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to User Input in Go"
	println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")

	// comma ok pattern

	input, _ := reader.ReadString('\n')
	fmt.Println("Hello", input)

	// number input by type conversion

	fmt.Println("Enter your age: ")

	input2, _ := reader.ReadString('\n')

	fmt.Println("You are", input2, "years old")

	age, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Next year you will be", age+1)
	}
	
}
