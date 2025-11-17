package main

import "fmt"

func main() {

	fmt.Println("Structs in Go Lang")

	// no inheritance in golan ; no super or parent class in go lang
	
	naiem := Usert{
		Name:   "Naiem",
		Email:  "naiem@example.com",
		Age:    30,
		Status: true,
	}

	fmt.Println(naiem)
	fmt.Printf("Details of naiem is: %+v\n", naiem)
	fmt.Printf("Name is: %v and Email is: %v", naiem.Name, naiem.Email)
}

type Usert struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
