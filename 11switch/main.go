package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	fmt.Println("Switch in go Lang")
	
	
	rand.Seed(time.Now().UnixNano())
	
	diceNumber := rand.Intn(6)+1
	
	fmt.Println("Dice number is:", diceNumber)
	
	
	switch diceNumber{
		case 1:
			fmt.Println("This is just One")
		case 2:
			fmt.Println("This is just Two")
		case 3:
			fmt.Println("This is just Three")
		case 4:
			fmt.Println("This is just Four")
		case 5:
			fmt.Println("This is just Five")
		case 6:
			fmt.Println("This is just Six")
			
		default:
			fmt.Println("This is the default Value")
	}
}