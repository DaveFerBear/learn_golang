package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{ name string }

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{ name string }

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	// Set up animal map.
	animalMap := make(map[string]string)
	animalMap["cow"] = "cow"
	animalMap["snake"] = "snake"
	animalMap["bird"] = "bird"

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSpace(input), " ")

		if len(s) != 3 {
			fmt.Println("Invalid input.")
			continue
		}

		if s[0] == "newanimal" {
			animalMap[s[1]] = s[2]
			fmt.Println("Created it!")
		} else if s[0] == "query" {
			var a Animal
			if animalMap[s[1]] == "cow" {
				a = Cow{name: s[1]}
			} else if animalMap[s[1]] == "bird" {
				a = Bird{name: s[1]}
			} else {
				a = Snake{name: s[1]}
			}
			switch s[2] {
			case "eat":
				a.Eat()
			case "speak":
				a.Speak()
			case "move":
				a.Move()
			}

		} else {
			fmt.Println("Invalid input.")
		}
	}
}
