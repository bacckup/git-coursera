/*
Write a program which prompts the user to enter
integers and stores the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should create an
empty integer slice of size (length) 3.
During each pass through the loop, the program
prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice,
sorts the slice, and prints the contents of the
slice in sorted order. The slice must grow in
size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var fSlice []int64
	reader := bufio.NewReader(os.Stdin)

	for i := 0; true; i++ {
		fmt.Print("Type a number: ")
		inputString, _ := reader.ReadString('\n')
		if strings.TrimSpace(inputString) == "X" {
			fmt.Println("Exiting a program.")
			break
		} else {
			inputNum, err := strconv.ParseInt(strings.TrimSpace(inputString), 10, 64)
			if err != nil {
				fmt.Println("Wrong input!")
			} else {
				fSlice = append(fSlice, inputNum)
				sort.Slice(fSlice, func(k, l int) bool {
					return fSlice[k] < fSlice[l]
				})
				fmt.Println(fSlice)
			}
		}
	}
}
