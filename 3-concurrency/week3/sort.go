package main

import (
	"fmt"
	"sync"
)

func Swap(s []int, i int) {
	s[i+1], s[i] = s[i], s[i+1]
}
func BubbleSort(s []int, wg *sync.WaitGroup) {
	fmt.Println("Sorting subarray: ", s)
	n := len(s)
	for j := 0; j < n-1; j++ {
		for i := 0; i < n-1; i++ {
			if s[i] > s[i+1] {
				Swap(s, i)
			}
		}
	}
	wg.Done()
}

func main() {
	const ARRAY_LENGTH int = 5

	// Make a slice
	arr := make([]int, ARRAY_LENGTH)

	var i int
	for d := ARRAY_LENGTH; d > 0; d-- {
		fmt.Printf("Enter %d more integers: ", d)
		fmt.Scanf("%d", &i)
		arr[ARRAY_LENGTH-d] = i
	}
	fmt.Println("Unsorted Array: ")
	fmt.Println(arr)

	const NUM_THREADS int = 4
	var wg sync.WaitGroup
	var sort_size = ARRAY_LENGTH / NUM_THREADS
	wg.Add(NUM_THREADS)
	for j := 0; j < NUM_THREADS; j++ {
		if j < NUM_THREADS-1 {
			go BubbleSort(arr[j*sort_size:(j+1)*sort_size], &wg)
		} else {
			go BubbleSort(arr[j*sort_size:], &wg) // last slice picks up the slack
		}
	}

	wg.Wait()

	fmt.Println("Sorted Array: ")
	fmt.Println(arr)
}
