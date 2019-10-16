package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
    consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a file name: ")
	fname, _ := consoleReader.ReadString('\n')
    fname = strings.TrimSpace(fname)

	file, _ := os.Open(fname)
	scanner := bufio.NewScanner(file)

	sl := make([]Name, 0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		n := Name{fname: string(s[0]), lname: string(s[1])}
		sl = append(sl, n)
	}

	for _, n := range sl {
		fmt.Printf("%s %s\n", string(n.fname[:]), string(n.lname[:]))
	}
}
