package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
	fmt.Println(e)
	
	// It’s important that you remember that := is used to declare the variable as well as assign a value to it.
	fmt.Println("This := is shorthand declaration")

    f := "apple"
	fmt.Println(f)
	// Also, a variable can’t be declared twice (not in the same scope anyway). If you try to run the following, you’ll get an error
	// f := "banana" 
	// If you uncomment the above line you will get this error: "no new variables on left side of :="
	g, f := "orange" ,"banana" 
	fmt.Println(g, f)
	// Although "f" is being used twice with :=, the compiler won’t complain the second time we use it, it’ll see that the
	// other variable, "g", is a new variable and allow :=. However, you can’t change the type of "f".

}