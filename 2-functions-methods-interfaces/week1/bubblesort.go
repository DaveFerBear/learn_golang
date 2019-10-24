package main

import (
	"fmt"
)

func Swap(s []int, i int) {
	s[i+1], s[i] = s[i], s[i+1]
}
func BubbleSort(s []int) {
	n := len(s)
	for j := 0; j < n-1; j++ {
		for i := 0; i < n-1; i++ {
			if s[i] > s[i+1] {
				Swap(s, i)
			}
		}
	}
}

func main() {
	SIZE := 10

	// Make a slice
	arr := make([]int, SIZE)

	// Read numbers into array
	var i int
	for d := SIZE; d > 0; d-- {
		fmt.Printf("Enter %d more integers: ", d)
		fmt.Scanf("%d", &i)
		arr[SIZE-d] = i
	}

	fmt.Print("The (unsorted) numbers you entered: ")
	fmt.Println(arr)

	BubbleSort(arr)

	fmt.Print("Sorted: ")
	fmt.Println(arr)
}
