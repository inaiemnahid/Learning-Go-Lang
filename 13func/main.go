package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Function in Go Lang")
	greeting()
	result := add(3,5)
	fmt.Println("Where is the result",result)
	
	
	
	newResult,message := proAdd(5, 4,3,3,3,3,3,3,3,3,3,3,2)
	fmt.Println("The total value is: ", newResult, "and the message is: ", message)
}


func greeting(){
	
	fmt.Println("Welcome again from go Lang function inside the function")
}

func add( val1 int, val2 int) int {
	return val1 + val2
}

func proAdd( values... int) (int, string) {
	sum :=0
	for val := range(values){
		sum = sum +val
	}
	return sum, "This is ALL OKay"
}