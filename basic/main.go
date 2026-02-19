package main

import "fmt"


func main() {
	
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go programming,", "Let's learn Go together.")

	var name string = "Alice"
	fmt.Println("My name is", name)

	age := 30
	fmt.Println("I am", age, "years old.")
	
	var a, b int = 5, 10
	fmt.Println("The sum of", a, "and", b, "is", a+b)

	const pi = 3.14
	fmt.Println("The value of pi is", pi)

}