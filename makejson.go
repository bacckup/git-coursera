/*
Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map
and add the name and address to the map
using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your name: ")
	nameInput, _ := reader.ReadString('\n')
	fmt.Print("Your address: ")
	addressInput, _ := reader.ReadString('\n')

	nameInput = strings.TrimSuffix(nameInput, "\n")
	addressInput = strings.TrimSuffix(addressInput, "\n")

	var person map[string]string = map[string]string{
		"name":    nameInput,
		"address": addressInput,
	}

	personJSON, _ := json.Marshal(person)

	fmt.Println("As JSON encoding: ", personJSON)
	fmt.Println("As String: ", string(personJSON))
}
