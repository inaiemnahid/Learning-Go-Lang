package main

import "fmt"

func main(){
	fmt.Println("Hello This is the example of defer")
	
	defer fmt.Println("This is line 1 here ")
	defer fmt.Println("This is line 2 here")
	fmt.Println("Hello World")
	Mydefer()
}


func Mydefer(){
	for i:= 0; i<5; i++{
	defer	fmt.Println("Print the value here", i)
	} 
}