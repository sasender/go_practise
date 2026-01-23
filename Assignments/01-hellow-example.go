package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	fmt.Print("Enter your Name:")
	fmt.Scan(&name)

	date := time.Now().Format("02-01-2006")
	fmt.Println("hellowo,", name, date)
}
