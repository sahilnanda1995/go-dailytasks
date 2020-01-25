package main

import "fmt"

func main() {

	// JavaScript has 3 loops: for loop, while loop, and a do while loop.
	// Go has only one type of loop, and it's the for loop. Don't let that 
	// deceive you though as the for loop in Go is very versatile and can emulate 
	// almost any type of loop.
    i := 1
    for i < 5 {
		fmt.Println(i)
		i = i+1
	}

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
	}
	
	// key value pairs
    kvs := map[string]string{
        "anime":    "OPM",
        "MC": "Saitama",
    }

    for key, value := range kvs {
        fmt.Println(key, value)
    }
}