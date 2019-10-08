package main

import (
  "fmt",
  "strings"
)

func main() {
  var input string
  fmt.Scan(&input)
  if strings.Index(input, "a") != 0 {
    fmt.Println("Not Found!)
  }
}

