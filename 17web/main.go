package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.nexifer.com/hello.html"

func main(){
	fmt.Println("This is the web in Go")
	
	response, err := http.Get(url)
	
	if err!= nil{
		panic(err)
	}
	
	fmt.Printf("This is the response %T\n", response)
	
	defer response.Body.Close()
	
	
	databytes, err := ioutil.ReadAll(response.Body)
	
	if err!=nil{
		panic(err)
	}
	
	fmt.Println(string(databytes))
}