package main

import (
	"fmt"
	"sync"
)

// Struct to hold food type and amount
type FoodItem struct {
	foodType string
	amount   int
}

func food(c chan FoodItem, foodtype string, foodamount int, wg *sync.WaitGroup) {
	// Send food type and amount through channel
	c <- FoodItem{foodType: foodtype, amount: foodamount}
	wg.Done() // Mark goroutine as done
}

func main() {
	var wg sync.WaitGroup
	var numOrders int

	fmt.Print("How many food orders do you want to place? ")
	fmt.Scanln(&numOrders)

	foodOrders := make([]FoodItem, numOrders)

	// Get food orders from user
	for i := 0; i < numOrders; i++ {
		var foodType string
		var amount int

		fmt.Printf("Order %d - Enter food type: ", i+1)
		fmt.Scanln(&foodType)

		fmt.Printf("Order %d - Enter amount: ", i+1)
		fmt.Scanln(&amount)

		foodOrders[i] = FoodItem{foodType: foodType, amount: amount}
	}

	// Create channel with buffer size equal to number of orders
	value := make(chan FoodItem, numOrders)

	// Add all goroutines to wait group
	wg.Add(numOrders)

	// Launch concurrent food requests
	for i := 0; i < numOrders; i++ {
		go food(value, foodOrders[i].foodType, foodOrders[i].amount, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Collect results
	var totalAmount int
	fmt.Println("\n--- Food Orders Results ---")
	for i := 0; i < numOrders; i++ {
		result := <-value
		totalAmount += result.amount
		fmt.Printf("Order %d: %s - Amount: %d\n", i+1, result.foodType, result.amount)
	}

	close(value)
	fmt.Printf("\nTotal amount of all food: %d\n", totalAmount)
}
