package main

import "fmt"

func main() {
	var x interface{} = "foo"

	var s string = x.(string)
	fmt.Println(s)     // "foo"

	s, ok := x.(string)
	fmt.Println(s, ok) // "foo true"

	n, ok := x.(int)
	fmt.Println(n, ok) // "0 false"

	// Following line will throw error: panic: interface conversion: interface {} is string, not int
	// n = x.(int)        // ILLEGAL
}