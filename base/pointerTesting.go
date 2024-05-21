package base

import (
	"fmt"
)

func main() {

	Testing()
}

func Testing() {
	var isstring string = "Testing value"
	var cup *string = &isstring

	fmt.Println("Isvalue", isstring, cup, *cup)

	isstring = "Changed value of var"
	fmt.Println("NextVal", isstring, cup, *cup)

	var strval2 string = "Apples are NOT too good"
	cup = &strval2

	fmt.Println("NextVal2", strval2, cup, &strval2, cup == &strval2, *cup)
}
