package main

import "fmt"

func main() {
	var a, b, res int
	var op string

	fmt.Scanln(&a)
	fmt.Scanln(&op)
	fmt.Scanln(&b)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(res)
}
