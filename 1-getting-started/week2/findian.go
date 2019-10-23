package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string \n")

	input, _ := consoleReader.ReadString('\n')

	lowerInput := strings.ToLower(strings.TrimSpace(input))

	a := strings.HasPrefix(lowerInput, "i")
	b := strings.HasSuffix(lowerInput, "n")
	c := strings.Index(lowerInput, "a") > 0

	fmt.Println("%t %t %t", a, b, c)

	if a && b && c {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
