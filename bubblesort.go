/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 10)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter up to 10 numbers (separated by space): ")
	scanner.Scan()
	InputToSlice(scanner.Text(), &sli, scanner)
	BubbleSort(&sli)
	PrintSlice(&sli)
}

func BubbleSort(slice *[]int) {
	for i := 1; i < len(*slice); i++ {
		for j := 1; j < len(*slice); j++ {
			if (*slice)[j] < (*slice)[j-1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice *[]int, index int) {
	tmp := (*slice)[index-1]
	(*slice)[index-1] = (*slice)[index]
	(*slice)[index] = tmp
}

func InputToSlice(input string, slice *[]int, sc *bufio.Scanner) {
	var count int
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("You entered not integer values or more than 10 elements")
			os.Exit(0)
		}
	}()

	for j, i := range strings.Fields(input) {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		} else {
			(*slice)[j] = num
		}
		count++
	}
	*slice = (*slice)[0:count]
}

func PrintSlice(slice *[]int) {
	fmt.Print("Result: ")
	for i := range *slice {
		if i == len(*slice)-1 {
			fmt.Printf("%d\n", (*slice)[i])
		} else {
			fmt.Printf("%d ", (*slice)[i])
		}
	}
}
