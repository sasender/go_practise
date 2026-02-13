package main

import (
	"errors"
	"fmt"
	"user" // Assuming the previous package is imported
)

func main() {
	err := user.CreateUser(nil)
	if err != nil {
		// 🔑 Using errors.Is to find a specific error value anywhere in the chain
		if errors.Is(err, user.ErrDatabaseTimeout) {
			// This is a known issue. We can tell the user to wait or try again. (Maybe a 503 HTTP status)
			fmt.Println("❌ Client: Database is unavailable. Retrying operation shortly.")
		} else {
			// Yikes, this is something else. Log the whole thing and maybe send a 500 error.
			fmt.Printf("⚠️ Client: Unhandled critical error: %v\n", err)
		}
	} else {
		fmt.Println("✅ Client: User created successfully.")
	}
}
