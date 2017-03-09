package main

import "fmt"

func main(){

var numbers[5]int				//initialize array,becomes [0,0,0,0,0]
numbers[2] = 100				//change one of them
some_numbers := numbers[0:4]	//create a slice of an array, [0:4]means 0,1,2,3
fmt.Println(some_numbers)
fmt.Println(len(numbers))

var scores []float64			//initialize a slice
scores = append(scores,1.1)		//append an float number
fmt.Println(scores)
scores[0] = 2.2
fmt.Println(scores)


}