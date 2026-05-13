package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	evenCount := 0
	oddCount := 0
	for _, num := range s {
		if num%2 == 0 {
			evenCount++

		} else {
			oddCount++
		}
	}
	fmt.Println("Even numbers count:", evenCount)
	fmt.Println("Odd numbers count:", oddCount)
}
