package main

import "fmt"

func main() {
	elements := make(map[string]int)
	elements["H"] = 1
	fmt.Println(elements["H"]+\n)

	// remove by key
	elements["O"] = 8
	delete(elements, "O")

	number, ok := elements["H"];
	fmt.Println(ok)
	fmt.Println(number) // 1
	
	// only do something with a element if it's in the map
	if number, ok := elements["O"]; ok {
		fmt.Println(number) // won't be printed
	}
	
	if number, ok := elements["H"]; ok {
		fmt.Println(ok)
		fmt.Println(number) // 1
	}
}