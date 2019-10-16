package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var count int
	sli := make([]int, 3, 10)
	stop := false
	input := ""

	for !stop {
		fmt.Print("Enter a number: ")
		fmt.Scan(&input)

		if input == "X" {
			stop = true
		} else {
			i, _ := strconv.Atoi(input)
			if count < 3 {
				for idx, n := range sli {
					if n == 0 {
						sli[idx] = i
						break
					}
				}
				count++
			} else {
				sli = append(sli, i)
			}
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}
}
