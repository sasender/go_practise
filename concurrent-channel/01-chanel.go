/*
package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "hello world"
	}()

	msg := <-messages

	fmt.Println(msg)

}
*/
package main

import (
	"fmt"
)

func mesg(mesg chan string) {
	mesg <- "hellow world"
	mesg <- "welcome to channel"
	close(mesg)
}

func main() {
	messages := make(chan string)

	go mesg(messages)
	for msg := range messages {
		fmt.Println(msg)
	}
}
