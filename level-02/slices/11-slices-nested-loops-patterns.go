package main

import "fmt"

func main() {
	rows := 4
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
