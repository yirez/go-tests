package main

import "fmt"

func main() {
	var myyarr [4]int
	myyarr[1] = 2
	myyarr[3] = 33

	for i, values := range myyarr {
		fmt.Println("index:", i, "val:", values)
	}
	myslices := []int{2, 3, 4, 4, 42, 2}

	for i, values := range myslices {
		fmt.Print("; index:", i, " - val:", values)
	}

	other_slice := []int{4, 322}
	myslices = append(myslices, 22)             //add values directly
	myslices = append(myslices, other_slice...) // add from another slice

	fmt.Println("\nmodified ")
	fmt.Println(myslices)
	// {2, 3, 4, 4, 42, 2, 22, 4, 322}
	//to remove a value/set of values from slice we create another slice from it
	//without that value. To remove 4 and 42 for example;

	fmt.Println("\nremoved vals")
	myslices = append(myslices[:3], myslices[5:]...)
	fmt.Println(myslices)
	fmt.Println(len(myslices))
	fmt.Println(cap(myslices))

	//maps
	my_map := map[string]int{
		"testkey":  11,
		"testkey2": 22,
	}
	fmt.Println("\nmap")
	fmt.Println(my_map)
	fmt.Println(my_map["testkey2"])

}
