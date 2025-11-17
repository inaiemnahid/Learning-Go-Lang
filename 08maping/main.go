package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	
	fmt.Println("Learning Map")
	
	languages := make(map[string]string)
	
	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["GO"] = "Go"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["JV"] = "Java"
	languages["C"] = "C"
	languages["CPP"] = "C++"
	
	fmt.Println(languages)
	
	fmt.Println("JS stands for :", languages["JS"])
	fmt.Println("TS Stands for :", languages["TS"])
	fmt.Println("PY Stands for", languages["PY"])
	fmt.Println("RB Stands for", languages["RB"])
	fmt.Println("JV Stands for", languages["JV"])
	fmt.Println("C Stands for", languages["C"])
	fmt.Println("CPP Stands for", languages["CPP"])
	fmt.Println("GO Stands for", languages["GO"])
	
	
	
	delete(languages, "PY")
	fmt.Println(languages)

	// loops are very interesting in go lang
	// 
	
	
	for key, value := range languages {
		fmt.Println(key, "stands for", value)
	}
	
	fmt.Println("Length of map is", len(languages))
}