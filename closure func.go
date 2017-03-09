package main

import "fmt"

func main() {

	number := 0

	/* It has to be a local variable like this.
	   You can't do `func increment(amount int) {` */
	increment := func(amount int) {
		number += amount
	}
	increment(1)
	increment(2)

	fmt.Println(number) // 3

}


/*
def run():

    def increment(amount):
        return number + amount

    number = 0
    number += increment(1)
    number += increment(2)
    print number  # 3

run()
*/