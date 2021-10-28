package vartrials

import "fmt"

global_short_decl_var := 42
var regular_global_var_decl = 22
var something_else string // declare with type

func vartrials() {
	some_var := 42                 // shorthand declare and assign
	some_other_var = 11            //undeclared and creates an error.
	var some_other_other_var = 122 // explicitly declared
	fmt.Println(some_var, some_other_var, some_other_other_var)

	this_var_is_valid := "some text here"
	also_this_string_is_valid := `"some" interesting 
	
	asignment
	`
	fmt.Println(this_var_is_valid)
	fmt.Println(also_this_string_is_valid)

	//scoped var, lives only within the curly braces
	{
		scopedVar := 1
		fmt.Println(scopedVar)
	}
	fmt.Println(scopedVar) //this throws an error because it doesn't exıst
}
