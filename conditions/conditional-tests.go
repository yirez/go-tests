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

	if i := 4; i == 4 { //local i, scoped within the function
		fmt.Println("i is 4")
	}

	switch i { //function i var which is 3
	case 1:
		fmt.Println("i 1")

	case 3, 52, 88:
		fmt.Println("i 3")
		fallthrough //this makes switch work like in c# and java
	case 4:
		fmt.Println("i 4")
	case 5:
		fmt.Println("i 5")
	}
}
