package main

import "fmt"

func main() {
	divide(1, 0)
}
func divide(x int, y int) {

	//this bit runs at the end regardless and
	//if we are recovering from something
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering. error is: %v \n", r)
		}
	}()

	result := x / y
	fmt.Printf("Division=%v \n", result)
}
