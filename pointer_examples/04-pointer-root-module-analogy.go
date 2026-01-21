package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== POINTERS LIKE ROOT MODULES IN TERRAFORM ===\n")

	// SCENARIO 1: Like Terraform Root Module - All "child modules" see changes
	fmt.Println("SCENARIO 1: Multiple Pointers (Like child modules referencing root module)")
	fmt.Println("---")

	// Root value (like root module variable)
	rootValue := 100

	// Multiple pointers referencing the same value (like child modules using root variable)
	pointer1 := &rootValue
	pointer2 := &rootValue
	pointer3 := &rootValue

	fmt.Printf("Initial value: %d\n", rootValue)
	fmt.Printf("pointer1 sees: %d\n", *pointer1)
	fmt.Printf("pointer2 sees: %d\n", *pointer2)
	fmt.Printf("pointer3 sees: %d\n", *pointer3)

	// Update through pointer1 (like changing root module variable)
	fmt.Println("\n--- Updating through pointer1 to 500 ---")
	*pointer1 = 500

	// All pointers see the change (like child modules seeing root module updates)
	fmt.Printf("After change - rootValue: %d\n", rootValue)
	fmt.Printf("pointer1 sees: %d\n", *pointer1)
	fmt.Printf("pointer2 sees: %d\n", *pointer2)
	fmt.Printf("pointer3 sees: %d\n", *pointer3)

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// SCENARIO 2: Copying values (NOT like root module - changes isolated)
	fmt.Println("SCENARIO 2: Copied Values (NOT like root modules - isolated)")
	fmt.Println("---")
	fmt.Println("WHY? Each variable gets its OWN memory address!")
	fmt.Println("originalConfig    → memory address 0x1000")
	fmt.Println("childModule1Config → memory address 0x2000 (DIFFERENT!)")
	fmt.Println("childModule2Config → memory address 0x3000 (DIFFERENT!)")
	fmt.Println("childModule3Config → memory address 0x4000 (DIFFERENT!)")
	fmt.Println("")

	originalConfig := 100

	// Copying values (like hardcoding values in each child module)
	childModule1Config := originalConfig
	childModule2Config := originalConfig
	childModule3Config := originalConfig

	fmt.Printf("Initial: original=%d, child1=%d, child2=%d, child3=%d\n",
		originalConfig, childModule1Config, childModule2Config, childModule3Config)

	// Update original
	fmt.Println("\n--- Updating originalConfig to 500 ---")
	fmt.Println("⚠️  This ONLY changes the value at 0x1000!")
	fmt.Println("Values at 0x2000, 0x3000, 0x4000 are UNTOUCHED!")
	originalConfig = 500

	// Changes NOT reflected in copies (unlike pointers)
	fmt.Printf("After change: original=%d, child1=%d, child2=%d, child3=%d\n",
		originalConfig, childModule1Config, childModule2Config, childModule3Config)

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	type DBConfig struct {
		Host    string
		Port    int
		MaxConn int
	}

	// Root module config
	rootDBConfig := DBConfig{
		Host:    "localhost",
		Port:    5432,
		MaxConn: 100,
	}

	// Child modules using pointers to root config
	authServiceConfig := &rootDBConfig
	paymentServiceConfig := &rootDBConfig
	userServiceConfig := &rootDBConfig

	fmt.Printf("Initial - authService MaxConn: %d\n", authServiceConfig.MaxConn)
	fmt.Printf("Initial - paymentService MaxConn: %d\n", paymentServiceConfig.MaxConn)
	fmt.Printf("Initial - userService MaxConn: %d\n", userServiceConfig.MaxConn)

	// Update through one pointer (like updating root module variable)
	fmt.Println("\n--- Updating MaxConn in root config to 200 ---")
	authServiceConfig.MaxConn = 200

	// All services see the update
	fmt.Printf("After change - authService MaxConn: %d\n", authServiceConfig.MaxConn)
	fmt.Printf("After change - paymentService MaxConn: %d\n", paymentServiceConfig.MaxConn)
	fmt.Printf("After change - userService MaxConn: %d\n", userServiceConfig.MaxConn)
	fmt.Printf("After change - original rootDBConfig MaxConn: %d\n", rootDBConfig.MaxConn)

	fmt.Println("\n✓ All child modules automatically receive the updated value!")
}
