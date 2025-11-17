package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to go loop")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	fmt.Println("Array Type 1")
	for d := 0; d < len(numbers); d++ {
		fmt.Println(numbers[d])
	}

	fmt.Println("Array Type 2")
	for i := range numbers {
		fmt.Println(numbers[i])
	}

	fmt.Println("Array Type 3")
	for index, num := range numbers {
		fmt.Printf("Index is %v and the value is %v\n", index, num)
	}

	fmt.Println("Comma OK syntax >> _, ok")
	for _, num := range numbers {
		fmt.Printf("This is the value %v\n", num)
	}

	fmt.Println("New type of loop")
	regularValue := 0
	for regularValue < 10 {
		fmt.Println(regularValue)
		regularValue++
	}

	fmt.Println("------------------------------")
	value := 1

	for value < 100 {
		if value == 10 {
			fmt.Println("This is 10")
		}

		value++

		if value == 15 {
			goto comehere
		}

		if value == 20 {
			fmt.Println("Value reached maximum value")
			break
		}
	}

comehere:
	fmt.Println("Here we go and all okay")

comeback:
	fmt.Println("Here we comeback again")
	return
	goto comeback
}
