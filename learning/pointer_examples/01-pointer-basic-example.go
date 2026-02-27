package main

import "fmt"

func main() {

	x := 10
	y := &x
	*y = 45
	fmt.Println("Value of x:", x)    // 10
	fmt.Println("Address of x:", &x) // memory address
	fmt.Println("Value of y:", y)    // same address as &x
	fmt.Println("Value pointed by y:", *y)
}
