package main

import "fmt"

func main() {
	a := 10
	b := 20
	if a > b {
		fmt.Printf("%d is greater than %d\n", a, b)
	} else if b > a {
		fmt.Printf("%d is greater than %d\n", b, a)
	} else {
		fmt.Printf("%d and %d are equal\n", a, b)

	}

}
