package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"strings"
)


func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a name: ")
	n, _ := consoleReader.ReadString('\n')
	n = strings.Trim(n, "\n")
	fmt.Print("Enter an address: ")
	a, _ := consoleReader.ReadString('\n')
    a = strings.Trim(a, "\n")
    data := map[string] string {"name": n, "address": a}
    j, _ := json.Marshal(data)

    fmt.Println(string(j))
}
