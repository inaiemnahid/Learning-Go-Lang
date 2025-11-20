package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main(){
	fmt.Println("This is how file works in Go")
	
	content := "Hello World this is the content of this file"
	
	file, errs := os.Create("./myfile.txt")
	
	if errs!= nil {
		panic(errs)
	}
	
	length, err := io.WriteString(file, content)

	if err!= nil{
		panic(err)
	}
	
	fmt.Println("The length is: ", length)
	file.Close()
	
	Readfile("./myfile.txt")
}



func Readfile(filename string){
	databyte , err := ioutil.ReadFile(filename)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(databyte))
}