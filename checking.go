package main

import (
	"fmt"
)

func Checker() {

	var currentState = func(name string) string {
		return name + "Apple"
	}
	fmt.Printf(currentState("The best fruit"))
}
