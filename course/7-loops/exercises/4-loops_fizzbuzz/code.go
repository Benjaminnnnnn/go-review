package main

import "fmt"

func fizzbuzz() {
	// ?
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println("buzz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
