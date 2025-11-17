package main

import "fmt"

func main(){
	fmt.Println("Learning Go if Else")
	
	
	loginCount := 24
	var result string
	
	if loginCount < 10{
		fmt.Println("Hey buddy this is less then 10")
	} else { 
		fmt.Println( "Hey buddy this is okay now")
	}
	fmt.Println("Here is your result", result)
}