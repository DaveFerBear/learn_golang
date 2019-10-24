package main

import (
	"fmt"
)

func genDisplaceFn(a float64, v float64, d float64) func(t float64) float64 {
	return func(t float64) float64{
	    dt := 0.5*a*t*t + v*t + d
	    return dt
	}
}

func main() {
	var a, v, d, t float64
	fmt.Printf("Enter acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Enter initial velocity: ")
	fmt.Scanf("%f", &v)
	fmt.Printf("Enter initial displacement: ")
	fmt.Scanf("%f", &d)
	fmt.Printf("Enter a time: ")
	fmt.Scanf("%f", &t)

	fn := genDisplaceFn(a, v, d)
	fmt.Println("The final position is: ")
	fmt.Println(fn(t))
}
