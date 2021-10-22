package main

import "fmt"

func main() {

	type new_type_who_dis string

	var this_works new_type_who_dis = "Hello?"

	fmt.Println(this_works)
	fmt.Printf("%T \n", this_works)

	var stringy string = "hi?"
	//this_works = stringy //this throws an error
	this_works = new_type_who_dis(stringy)
}
