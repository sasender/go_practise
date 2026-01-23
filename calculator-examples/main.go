package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Simple Calculator App ===")
	fmt.Println("Select operation: 1=Add, 2=Subtract, 3=Multiply, 4=Divide")
	fmt.Println()

	for {
		// Get first number
		num1 := getNumber("Enter first number (or 'exit' to quit): ")
		if num1 == nil {
			fmt.Println("Goodbye!")
			break
		}

		// Get operator
		operator := getOperator("Enter operation (1=Add, 2=Subtract, 3=Multiply, 4=Divide): ")
		if operator == "" {
			continue
		}

		// Get second number
		num2 := getNumber("Enter second number: ")
		if num2 == nil {
			continue
		}

		// Perform calculation
		result, err := calculate(*num1, *num2, operator)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			continue
		}

		opSymbol := getOperatorSymbol(operator)
		fmt.Printf("Result: %.2f %s %.2f = %.2f\n\n", *num1, opSymbol, *num2, result)
	}
}

// getNumber reads a number from user input
func getNumber(prompt string) *float64 {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Check if user wants to exit
	if strings.ToLower(input) == "exit" {
		return nil
	}

	// Convert string to float64
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid number! Please enter a valid number.")
		return getNumber(prompt)
	}

	return &num
}

// getOperator reads an operator from user input (1, 2, 3, 4)
func getOperator(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Validate operator choice
	validOps := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
	}

	if !validOps[input] {
		fmt.Println("Invalid operation! Please use 1, 2, 3, or 4")
		return getOperator(prompt)
	}

	return input
}

// calculate performs the arithmetic operation
func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "1":
		return num1 + num2, nil
	case "2":
		return num1 - num2, nil
	case "3":
		return num1 * num2, nil
	case "4":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}

// getOperatorSymbol returns the symbol for the operation number
func getOperatorSymbol(op string) string {
	switch op {
	case "1":
		return "+"
	case "2":
		return "-"
	case "3":
		return "*"
	case "4":
		return "/"
	default:
		return "?"
	}
}
