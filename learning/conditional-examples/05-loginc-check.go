package main

import (
	"fmt"
)

func main() {
	username := "admin"
	password := "1234"
	if username == "admin" && password == "1234" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Invalid")
	}
}
