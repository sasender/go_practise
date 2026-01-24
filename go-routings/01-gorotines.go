package main

import (
	"fmt"
	"time"
)

func dance() {
	fmt.Println("he is dancing")
}

func sing() {
	fmt.Println("he is singing")
}

func main() {
	go dance()
	go sing()
	time.Sleep(1 * time.Second) // wait for goroutines to finish
	fmt.Println("he is singing")

}
