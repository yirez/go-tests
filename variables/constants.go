package main

import "fmt"

const (
	one          = "ONE"
	other        = iota
	other_second = iota
	how
)

func main() {
	fmt.Println(one)
	fmt.Println(other)
	fmt.Println(other_second)
	fmt.Println(how)
}
