package main

import "fmt"

var x int

func main() {
	x = 11
	fmt.Println(x)
	fmt.Println(&x) //returns the addr of value
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)
	checkWhatsInAddr(&x)

}
func checkWhatsInAddr(addressGiven *int) {
	fmt.Println(*addressGiven)
	changedVal := *addressGiven + 10
	*addressGiven = 44
	fmt.Println(changedVal)
	fmt.Println(x)
}
