package main

import (
	"errors"
	"fmt"
)

// checkNumber function checks if a number is positive or negative.
// If the number is negative, it returns an error.
func checkNumber(a int, b int) (string, error) {
	if a < 0 {
		// Return an error when the number is negative
		return "", errors.New("number is negative")
	}
	if b < 0 {
		// Return an error when the number is negative
		return "", errors.New("number is negative")
	}
	if a == 0 && b == 0 {
		// Return a success message if both numbers are zero
		return "both numbers are zero", nil
	}
	// Return a success message if both numbers are positive
	return "both numbers are positive", nil
}

func main() {
	// Calling checkNumber with a negative value (-5)
	var (
		a, b int
	)
	fmt.Print("Enter the values a, b: ")
	fmt.Scanf("%d %d", &a, &b)

	result, err := checkNumber(a, b)

	if err != nil {
		// If an error is returned, print the error message
		fmt.Println("Error:", err) // Output: Error: number is negative
	} else {
		// If no error, print the success message
		fmt.Println(result)
	}
}
