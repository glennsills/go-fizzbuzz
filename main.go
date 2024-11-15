package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a command line flag
	maxvalue := flag.Int("maxvalue", 15, "the maximum integeter to test")

	// Parse the command line flags
	flag.Parse()

	// Use the flag value
	fmt.Println("The maximum integer to test is", *maxvalue)

	for i := 1; i <= *maxvalue; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")

		case i%3 == 0:
			fmt.Println("Fizz")

		case i%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(i)
		}
	}
}
