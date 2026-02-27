package main

import "fmt"

func multiple(c chan int, someValue int) {
	c <- someValue * 2
}

func main() {
	var num1, num2 int

	// Ask user for input
	fmt.Print("Enter first value (integer): ")
	_, err1 := fmt.Scanln(&num1)

	fmt.Print("Enter second value (integer): ")
	_, err2 := fmt.Scanln(&num2)

	// Handle input errors
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please enter valid integers")
		return
	}

	value := make(chan int)
	go multiple(value, num1)
	go multiple(value, num2)

	v1 := <-value
	v2 := <-value

	fmt.Println("Results:", v1, v2)
}
