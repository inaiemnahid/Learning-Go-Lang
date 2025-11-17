package main

import "fmt"

func main(){
	fmt.Println("Welcome to a Go Pointers")
	// A pointer is a variable that stores the memory address of another variable.
	// It allows you to indirectly access and modify the value of a variable.
	// It is useful for passing large data structures to functions without copying them.
	
	
	var ptr *int
	// fmt.Println("Value of pointer is ", ptr)
	// 
	myNumber := 23
	
	ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value in pointer is ", *ptr)
	
	*ptr = *ptr  * 2
	fmt.Println("New value is: ", myNumber)
}
