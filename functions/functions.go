package main

import (
	"errors"
	"fmt"
)

func main() {
	parameterlessFunc()
	funcWithOneParameter(22)
	funcWithOneSliceParameterLikeFmtPrintln("hello", "several", "itemsToPrint")
	data, errorVal := funcWithOneParameterAndSeveralReturnVals(44)
	fmt.Println(data, "\n", errorVal.Error())

	func(par int) {
		fmt.Println("this is from an anon func, pass par:", par)
	}(44)

	//using functions like a var
	someFunc := func(par int) {
		fmt.Println("this is from a func as val, pass par:", par)
	}
	someFunc(4123)

	total := add(1, 2, 3, 4)
	fmt.Println(total)

	total = addEven(add, 33, 1, 2, 3, 4)
	fmt.Println(total)

}

func parameterlessFunc() { //this { has to be here.
	//	Go impilicity places ; after statements
	fmt.Println("Hello from parameterless")
}

func funcWithOneParameter(par int) {
	fmt.Println("Hello with parameter:", par)
}

func funcWithOneSliceParameterLikeFmtPrintln(par ...string) {
	fmt.Println("Hello with parameter:", par)
}

func funcWithOneParameterAndSeveralReturnVals(par int) (data string, genericError error) {
	fmt.Println("Hello with parameter:", par)
	return "This is returned from func", errors.New("this is an error")
}

func add(valListToAdd ...int) int {
	total := 0
	for _, val := range valListToAdd {
		total += val
	}
	fmt.Println("total:", total)
	return total
}

func addEven(addfunc func(valListToAdd ...int) int, valListToAdd ...int) int {
	var evenNums []int
	for _, val := range valListToAdd {
		if val%2 == 0 {
			evenNums = append(evenNums, val)
		}
	}
	return addfunc(evenNums...)
}
