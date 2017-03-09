package main

import "fmt"

func main(){
	fmt.Println("test Println") //change line directly
	
	fmt.Print("test Print\n")
	
	fmt.Printf("Name:%s,Age:%d\n","Peter",35)
	
	//use ` not '
	fmt.Println(`This is 
	a multi-line string.
	`)	
	
	//O'word Another "word" Last word.
	fmt.Println(
		"O'word " +
			"Another \"word\" " +
			"Last word.")   
}	