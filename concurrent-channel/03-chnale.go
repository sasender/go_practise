package main

import "fmt"

// Struct to hold food type and amount
type FoodItem struct {
	foodType string
	amount   int
}

func food(c chan FoodItem, foodtype string, foodamount int) {
	// Send food type and amount through channel
	c <- FoodItem{foodType: foodtype, amount: foodamount}
}

func main() {
	var food1, food2 string
	var amount1, amount2 int

	fmt.Print("Enter the first food type: ")
	_, err1 := fmt.Scanln(&food1)

	fmt.Print("Enter the first food amount: ")
	_, errA1 := fmt.Scanln(&amount1)

	fmt.Print("Enter the second food type: ")
	_, err2 := fmt.Scanln(&food2)

	fmt.Print("Enter the second food amount: ")
	_, errA2 := fmt.Scanln(&amount2)

	if err1 != nil || err2 != nil || errA1 != nil || errA2 != nil {
		fmt.Println("Error: Please enter valid inputs")
		return
	}

	value := make(chan FoodItem, 2)
	go food(value, food1, amount1)
	go food(value, food2, amount2)

	v1 := <-value
	v2 := <-value

	totalAmount := v1.amount + v2.amount

	fmt.Println("\nResults:")
	fmt.Println("Food 1:", v1.foodType, "Amount:", v1.amount)
	fmt.Println("Food 2:", v2.foodType, "Amount:", v2.amount)
	fmt.Println("Total amount of food:", totalAmount)
}
