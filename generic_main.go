package main

import (
	"fmt"

	"somedomain.com/go-tests/modules1"
)

func main() {
	n, _ := fmt.Println(modules1.HelloWithQuotes())
	fmt.Println(n)
}
