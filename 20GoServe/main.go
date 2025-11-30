package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main(){
	fmt.Println("Welcome to golang")
	webreq()
	postreq()
}

func webreq(){
	const myurl = "http://localhost:6677/"
	
	response,err := http.Get(myurl)
	if err!= nil{
		panic(err)
	}
	
	
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)
	
	res, err := ioutil.ReadAll(response.Body)
	
	fmt.Println(string(res))
}


func postreq(){
	const myurl ="http://localhost:6677/api/post"

	requestBody := strings.NewReader(`
		{
		"hello":"Hello",
		"Price": "Price"
		}
		`)
	
	response ,err := http.Post(myurl, "application/json", requestBody)
	if err != nil{
		
	}
	defer response.Body.Close()
	
	content,_ := ioutil.ReadAll(response.Body)
	
	fmt.Println(string(content))
	
}