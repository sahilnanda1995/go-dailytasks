package main

import "fmt"

func main()  {
	// A type switch performs several type assertions in series and runs the first case with a matching type.
	var x interface{} = "foo"

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")            // here v has type interface{}
	case int: 
		fmt.Println("x is", v)             // here v has type int
	case bool, string:
		fmt.Println("x is bool or string") // here v has type interface{}
	default:
		fmt.Println("type unknown")        // here v has type interface{}
	}
}