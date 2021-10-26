package main

import (
	"fmt"
)

func main() {

	type custom_type struct {
		field1 string
		field2 int
		field3 string
	}

	type child_type struct {
		uniq_field string
		custom_type
		master_prop custom_type
	}

	var so_custom custom_type = custom_type{
		field1: "this is a val",
		field2: 42,
		field3: "other prop",
	}

	var children child_type = child_type{
		uniq_field: "wow, much unique",
		custom_type: custom_type{
			field1: " a value",
			field2: 51,
			field3: "a prop",
		},
		master_prop: custom_type{
			field1: "1",
			field2: 33,
			field3: "44",
		},
	}

	fmt.Println(so_custom)
	fmt.Println(children)

	anon_val := struct {
		some_field    string
		another_field int
	}{
		some_field:    "well, this is awkward",
		another_field: 22,
	}
	fmt.Println(anon_val)

}
