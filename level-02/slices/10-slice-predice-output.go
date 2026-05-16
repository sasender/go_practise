package main

import "fmt"

func main() {
	s := make([]int, 2, 4)

	fmt.Println(len(s), cap(s))

	s = append(s, 10, 20)

	fmt.Println(len(s), cap(s))
}
