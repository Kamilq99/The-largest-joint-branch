package main

import (
	"fmt"
)

func gcd(a int, b int) int {

	if b == 0 {
		return a
	} else if a == 0 {
		return b
	} else {
		for a != b {
			if a > b {
				a = a - b
			} else {
				b = b - a
			}
		}
	}

	return a
}

func main() {

	var a int
	var b int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)

	result := gcd(a, b)

	fmt.Printf("Greatest common divisor of %d and %d = %d", a, b, result)
}
