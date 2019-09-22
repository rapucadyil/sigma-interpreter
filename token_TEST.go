package main

import "fmt"

func main() {
	c := "if"
	function := "proc"
	other_function := "function"
	d := "math"

	fmt.Printf("Look Up for %v = %v\n", c, look_up_identifier(c))
	fmt.Printf("Look Up for %v = %v\n", function, look_up_identifier(function))
	fmt.Printf("Look Up for %v = %v\n", other_function, look_up_identifier(other_function))
	fmt.Printf("Look Up for %v = %v\n", d, look_up_identifier(d))
}
