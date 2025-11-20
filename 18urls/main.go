package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.nexifer.com:3000/hello.html?hello=man&this=343"

func main(){
	fmt.Println("Welcome to the URL verse of Go Lang")
	
	result, _ := url.Parse(myurl)
	
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	
	qparams := result.Query()
	fmt.Printf("This is the query: %T\n", qparams)
	
	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "www.nexifer.com",
		Path: "about",
		RawPath: "/hello",
		RawQuery: "man=39030",
	}
	
	newUrl := partsofUrl.String()
	
	fmt.Println("This is new url: ", newUrl)
}