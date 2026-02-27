package main

import "fmt"

// Struct to hold food type and amount
type FoodItem struct {
	foodType string
	amount   int
}

// Map to store food types with their amounts
var foodInventory = map[string]int{
	"pizza":   5,
	"burger":  10,
	"pasta":   8,
	"salad":   15,
	"drink":   20,
	"beer":    180,
	"biryani": 135,
	"idli":    50,
	"poori":   30,
	"dosha":   25,
}

func food(c chan FoodItem, foodtype string) {
	// Check if food type exists in map
	if amount, exists := foodInventory[foodtype]; exists {
		c <- FoodItem{foodType: foodtype, amount: amount}
	} else {
		c <- FoodItem{foodType: foodtype, amount: 0}
	}
}

func main() {
	var food1, food2, food3, food4 string
	fmt.Println("Available foods:", foodInventory)

	fmt.Print("Enter the first food type: ")
	_, err1 := fmt.Scanln(&food1)

	fmt.Print("Enter the second food type: ")
	_, err2 := fmt.Scanln(&food2)

	fmt.Print("Enter the third food type: ")
	_, err3 := fmt.Scanln(&food3)

	fmt.Print("Enter the fourth food type: ")
	_, err4 := fmt.Scanln(&food4)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Error: Please enter valid strings")
		return
	}

	value := make(chan FoodItem, 4)
	go food(value, food1)
	go food(value, food2)
	go food(value, food3)
	go food(value, food4)

	v1 := <-value
	v2 := <-value
	v3 := <-value
	v4 := <-value

	totalAmount := v1.amount + v2.amount + v3.amount + v4.amount

	fmt.Println("Results:", v1.foodType, v1.amount, "and", v2.foodType, v2.amount, "and", v3.foodType, v3.amount, "and", v4.foodType, v4.amount)
	fmt.Println("Total amount of food:", totalAmount)
}
