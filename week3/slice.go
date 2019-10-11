package main

import (
    "fmt"
    "strconv"
    )

func main() {
	var input string
	stop := false
	sum := 0

	for !stop {
	    fmt.Print("Enter a number: ")
	    fmt.Scan(&input)
	    if input == "X" {
	        stop = true
	    } else {
            i, _ := strconv.Atoi(input)
            sum += i
            fmt.Printf("New Sum: %d \n", sum)
	    }
	}
}
