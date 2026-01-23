package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	fmt.Print("Enter your Name:")
	//after we given the input its automatically stored in name variable by helping of scan
	fmt.Scan(&name)

	date := time.Now().Format("02-01-2006")
	fmt.Println("hellowo,", name, date)
}
