package main

import (
	"fmt"
)

type someCustomTypeName struct {
	field1 string
	field2 int
}

type someOtherCustomTypeName struct {
	field3 string
	field4 int
}

//a value can be of multiple types
//since we attached someCustomTypeName to typeActionFunc,
// through the interface below, they are also of type someInterfaceName
type someInterfaceName interface {
	typeActionFunc()
}

func main() {
	normalFunc()
	defer closeFunc()
	normalActionFunc()

	someData := someCustomTypeName{
		field1: "ook",
		field2: 22,
	}

	someOtherData := someOtherCustomTypeName{
		field3: "not ok",
		field4: 32,
	}
	someData.typeActionFunc()
	someOtherData.typeActionFunc()
	interfaceFunc(someData)
	interfaceFunc(someOtherData)
	interfaceFunc(&someData) //works with pointer values as well.
} // defer keyword waits until the end of parent func to run

func normalFunc() {
	fmt.Println("opening stuff")
}
func normalActionFunc() {
	fmt.Println("doing stuff")
}

//this func attaches itself to all someCustomTypeName types
func (customval someCustomTypeName) typeActionFunc() {
	fmt.Printf("type:%T - ", customval)
	fmt.Println("field1: ", customval.field1, "\n", "field2: ", customval.field2)
}

//this func attaches itself to all someOtherCustomTypeName types
func (customval someOtherCustomTypeName) typeActionFunc() {
	fmt.Printf("type:%T - ", customval)
	fmt.Println("field3: ", customval.field3, "\n", "field4: ", customval.field4)
}

//this func is used via the interface connections
func interfaceFunc(interfacedType someInterfaceName) {
	fmt.Printf("In interfaceFunc with type:%T\n", interfacedType)

	switch interfacedType.(type) {
	case someCustomTypeName:
		fmt.Println("field1: ", interfacedType.(someCustomTypeName).field1, "\n", "field2: ", interfacedType.(someCustomTypeName).field2)
	case someOtherCustomTypeName:
		fmt.Println("field3: ", interfacedType.(someOtherCustomTypeName).field3, "\n", "field4: ", interfacedType.(someOtherCustomTypeName).field4)
	case *someCustomTypeName:
		fmt.Println("field1: ", interfacedType.(*someCustomTypeName).field1, "\n", "field2: ", interfacedType.(*someCustomTypeName).field2)
	default:
		fmt.Println(interfacedType)
	}
}

func closeFunc() {
	fmt.Println("closing stuff")
}
