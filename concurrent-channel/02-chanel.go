package main

import "fmt"

func multiple(c chan int, someValue int) {
	c <- someValue * 2
}

func main() {
	value := make(chan int)
	go multiple(value, 10)
	go multiple(value, 5)

	v1 := <-value
	v2 := <-value

	fmt.Println(v1, v2)
}
