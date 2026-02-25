package main

import (
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// enter input values for a and b
	var a, b int
	fmt.Print("Enter two integers (a b): ")
	fmt.Scanf("%d %d", &a, &b)

	value, err := divide(a, b)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(" Result:", value)
}
