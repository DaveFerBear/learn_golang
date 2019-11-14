/*

Q: What is a race condition?
A race condition is a situation where two threads try to access the same memory at the same time.
The results of these kinds of programs are often non-deterministic, as the order in which the
memory is accessed is determined by the thread scheduler, not the programmer.

The below program demonstrates a race condition because the output depends entirely on the thread scheduler.
If the program were not run concurrently, the output would be the numbers, followed by the letters.
Since these run concurrently, the actual output is a non-deterministic combination of both.

The program also demonstrates a second race condition, where sometimes the number 5 or letter e are printed,
and other times they are not.  This is because the program occasionally finishes the main thread time.Sleep(1000)
before the threads fully print, and other times they do not.
*/

package main

import (  
    "fmt"
    "time"
)

func print_numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(200 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func print_alphabet() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(200 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    
    fmt.Println("The following program will print the numbers from 1->5 and letters from a->e.")
    go print_numbers()
    go print_alphabet()
    time.Sleep(1000 * time.Millisecond)
}

