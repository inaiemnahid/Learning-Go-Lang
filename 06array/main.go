package main

import "fmt"

func main(){
	fmt.Println(("Welcome to Array in golang"))
	var fuitList [4]string
	
	fuitList[0] = "Apple"
	fuitList[1] = "Banana"
	fuitList[2] = "Orange"
	fuitList[3] = "Grapes"
	
	fmt.Println(fuitList)
	fmt.Println("Fruite Length: ", len(fuitList))
}