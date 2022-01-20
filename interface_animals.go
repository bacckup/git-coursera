/*
The same as "animals.go" but...
Define an interface type called Animal which describes the
methods of an animal. Specifically, the Animal interface
should contain the methods Eat(), Move(), and Speak(), which
take no arguments and return no values. The Eat() method
should print the animal’s food, the Move() method should
print the animal’s locomotion, and the Speak() method should
print the animal’s spoken sound. Define three types Cow,
Bird, and Snake. For each of these three types, define
methods Eat(), Move(), and Speak() so that the types Cow, Bird,
and Snake all satisfy the Animal interface. When the user
creates an animal, create an object of the appropriate type. Your program
should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	animals := make(map[string]AnimalInterface)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		if len(words) == 3 {
			if words[0] == "query" {
				if _, ok := animals[words[1]]; !ok {
					fmt.Printf("There is no animal with name \"%s\"\n", words[1])
				} else {
					switch words[2] {
					case "eat":
						animals[words[1]].Eat()
					case "move":
						animals[words[1]].Move()
					case "speak":
						animals[words[1]].Speak()
					default:
						fmt.Println("Wrong request")
					}
				}
			} else if words[0] == "newanimal" {
				switch words[2] {
				case "cow":
					animals[words[1]] = Cow{"grass", "walk", "moo"}
					fmt.Println("Created it!")
				case "bird":
					animals[words[1]] = Bird{"worms", "fly", "peep"}
					fmt.Println("Created it!")
				case "snake":
					animals[words[1]] = Snake{"mice", "slither", "tsss"}
					fmt.Println("Created it!")
				default:
					fmt.Println("Wrong type")
				}
			} else {
				fmt.Println("Wrong request")
			}
		} else {
			fmt.Println("Wrong request")
		}
	}
}
