package main

import "fmt"

func main() {
	var input float64
	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&input)
	fmt.Printf("Truncated Number: %d \n", int(input))
}
