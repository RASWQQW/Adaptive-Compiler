package main

import "fmt"

func main() {

	var a []int = []int{1, 2, 3}
	var b []int = []int{3, 4, 5}
	var c []int = append(a, b...)

	fmt.Println(a)
	fmt.Println("Len", c[0:len(c)])
	// var keys map[string]interface{}
	// Checker("apple", keys, c...)
}

func Checker(val string, keyValues map[string]interface{}, vals ...int) {
	fmt.Println(vals)
}
