package main

import (
	"fmt"
)


func swap(sli []int, i int, j int) {
    sli[i], sli[j] = sli[j], sli[i]
}

func bubbleSort(sli []int) {
    for i, a := range sli {
        for j, b := range sli {
            if i!=j && a < b {
                swap(sli, i, j)
            }
        }
    }
}

SIZE := 10

func main() {
    // Make a slice
    arr := make([]int, SIZE)

    // Read numbers into array
    var i int
    for d:=SIZE; d>0; d-- {
        fmt.Printf("Enter %d more integers: ", d)
	    fmt.Scanf("%d", &i)
	    arr[SIZE-d] = i
    }

    fmt.Print("The (unsorted) numbers you entered: ")
    fmt.Println(arr)

    bubbleSort(arr)

    fmt.Print("Sorted: ")
    fmt.Println(arr)
}