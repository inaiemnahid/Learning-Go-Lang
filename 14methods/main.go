package main

import "fmt"


func main(){
	
	naiem := User{"Naiem", 20, "naiem@gmail.com", 10}
	
	fmt.Println("The one age is: ", naiem.oneAge)
	naiem.GetStatus()
	
	naiem.UpdateAge(30)
	fmt.Println("Again the age of naiem is: ", naiem.oneAge)
}

type User struct{
	Name string
	Age int
	Email string
	oneAge int
}

func (u User) GetStatus(){
	fmt.Println("The use age is: ", u.oneAge)
}


func (u User) UpdateAge( age int){
	u.oneAge = age 
	
	fmt.Println("New age of user is", u.oneAge)
	
}