package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const ARRAY_LENGTH int = 20
	var arr [ARRAY_LENGTH]int
	for i := 0; i < ARRAY_LENGTH; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Print(arr)
}
