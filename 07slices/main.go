package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Hello World!")
	
	var fruits = []string{"Apple", "Banana", "Cherry", "Grapes"}
	fmt.Println(fruits)
	fmt.Println("Fruit Length:", len(fruits))
	fmt.Printf("Type of fruitlist is %T\n", fruits)

	fruitList := append(fruits, "Mango", "Orange")
	
	fmt.Println(fruitList)
	fmt.Println("Fruit Length:", len(fruitList))
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println("Fruit Length:", len(fruitList))
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	
	
	highScores := []int{100, 200, 300, 400}
	fmt.Println(highScores)
	
	anotherhigheScore := make([]int, 4)
	
	anotherhigheScore[0] = 100
	anotherhigheScore[1] = 200
	anotherhigheScore[2] = 300
	anotherhigheScore[3] = 400
	
	anotherhigheScore = append(highScores, 343, 343, 343)
	fmt.Println(anotherhigheScore)
	fmt.Println(sort.IntsAreSorted((anotherhigheScore)))
	
	sort.Ints(anotherhigheScore)
	fmt.Println(anotherhigheScore)
	
	var courses = []string{"reactjs", "javascript", "Swift", "Vue", "Golang"}
	fmt.Println(courses)
	
	var index int = 2
	
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
