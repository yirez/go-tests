package main

import "fmt"

func main() {
	var i int = 0

	fmt.Println("loop1")
	for i = 0; i < 2; i++ {
		fmt.Println(i)
	}
	//or
	fmt.Println("loop2")
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//or
	fmt.Println("loop3")
	for {
		//stuff
		if i < 5 {
			fmt.Println("breaking")
			break
		}
		fmt.Println("continuing")
		continue
	}
}
