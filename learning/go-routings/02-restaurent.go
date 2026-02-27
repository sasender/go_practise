package main

import (
	"fmt"
	"time"
)

func menu(done chan bool) {
	fmt.Println("1. Placed menu on every table")
	time.Sleep(500 * time.Millisecond)
	done <- true // Signal completion
}

func orderFood(done chan bool) {
	fmt.Println("2. Taking order from customer")
	time.Sleep(500 * time.Millisecond)
	done <- true
}

func prepareFood(done chan bool) {
	fmt.Println("3. Preparing food in kitchen")
	time.Sleep(500 * time.Millisecond)
	done <- true
}

func serverFood(done chan bool) {
	fmt.Println("4. Serving food to customer")
	time.Sleep(500 * time.Millisecond)
	done <- true
}

func payment(done chan bool) {
	fmt.Println("5. Payment done by customer")
	time.Sleep(500 * time.Millisecond)
	done <- true
}

func main() {
	done := make(chan bool)

	// Run each step and wait for it to complete before next
	go menu(done)
	<-done // Wait for menu to finish

	go orderFood(done)
	<-done // Wait for order to finish

	go prepareFood(done)
	<-done // Wait for preparation to finish

	go serverFood(done)
	<-done // Wait for serving to finish

	go payment(done)
	<-done // Wait for payment to finish

	fmt.Println("\n✅ Restaurant cycle complete!")
}
