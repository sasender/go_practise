package main

import "fmt"

// Interface: defines a behavior
type Payment interface {
	Pay(amount int)
}

// Concrete type 1
type UPI struct {
	app string
}

func (u UPI) Pay(amount int) {
	fmt.Printf("UPI payment using %s: %d\n", u.app, amount)
}

// Concrete type 2
type Card struct {
	bank string
}

func (c Card) Pay(amount int) {
	fmt.Printf("Card payment using %s: %d\n", c.bank, amount)
}

func main() {
	var p Payment                      // Interface variable, currently nil
	fmt.Printf("p initially: %v\n", p) // nil

	// Assign a UPI struct
	u := UPI{"GPay"}
	p = u
	fmt.Printf("p after assigning UPI: %T, %v\n", p, p) // Type = UPI, Value = {GPay}
	p.Pay(100)                                          // Calls UPI.Pay

	// Assign a Card struct
	c := Card{"HDFC"}
	p = c
	fmt.Printf("p after assigning Card: %T, %v\n", p, p) // Type = Card, Value = {HDFC}
	p.Pay(500)                                           // Calls Card.Pay

	// Assign a nil pointer of UPI
	var uNil *UPI = nil
	p = uNil
	fmt.Printf("p after assigning nil UPI pointer: %T, %v\n", p, p) // Type = *UPI, Value = <nil>
	// Interface is NOT nil because type exists
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is NOT nil")
	}
	if u, ok := p.(*UPI); ok {
		if u == nil {
			fmt.Println("underlying pointer is nil")
		}
	}

}
