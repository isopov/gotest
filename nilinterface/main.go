package main

import "fmt"

func main() {
	fmt.Print(foo())
	fmt.Print(", ")
	fmt.Println(foo() == nil)
	fmt.Print(bar())
	fmt.Print(", ")
	fmt.Println(bar() == nil)
	fmt.Print(baz())
	fmt.Print(", ")
	fmt.Println(baz() == nil)
}

func foo() interface{} {
	return nil
}
func bar() interface{} {
	var foo *error
	return foo
}
func baz() interface{} {
	var foo error
	return foo
}
