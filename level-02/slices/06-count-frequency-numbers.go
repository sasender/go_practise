package main

import "fmt"

func main() {
	s := []int{1, 2, 2, 3, 3, 3}
	frequency := make(map[int]int)
	for _, value := range s {
		frequency[value]++
	}
	fmt.Println(frequency)
}
