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

// Merges an array with 2 sorted, subarrays.
func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func main() {
	const ARRAY_LENGTH int = 9

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

	// SORT
	var wg sync.WaitGroup
	var sort_size = ARRAY_LENGTH / 4
	wg.Add(4)
	for j := 0; j < 4; j++ {
		if j < 3 {
			go BubbleSort(arr[j*sort_size:(j+1)*sort_size], &wg)
		} else {
			go BubbleSort(arr[j*sort_size:], &wg) // last slice picks up the slack
		}
	}
	wg.Wait()

	// MERGE
	l := merge(arr[:sort_size], arr[sort_size:sort_size*2])
	r := merge(arr[sort_size*2:sort_size*3], arr[sort_size*3:])

    a := append(l, r...)
	b := merge(a[:sort_size*2], a[sort_size*2:])

	fmt.Println("Sorted Array: ")
	fmt.Println(b)
}
