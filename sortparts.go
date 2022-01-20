/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print
the subarray that it will sort. When sorting is complete,
the main goroutine should print the entire sorted list.
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

func sortPart(sli []int, c chan []int) {
	fmt.Println("Some goroutine is going to sort:", sli)
	sort.Ints(sli)
	c <- sli
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sli := make([]int, 0)
	sortedSli := make([]int, len(sli))
	ch := make(chan []int)

	readInput(&sli, scanner)
	partLen := len(sli) / 4

	go sortPart(sli[:partLen], ch)
	go sortPart(sli[partLen:partLen*2], ch)
	go sortPart(sli[partLen*2:partLen*3], ch)
	go sortPart(sli[partLen*3:], ch)

	for i := 0; i < 4; i++ {
		part := <-ch
		sortedSli = append(sortedSli, part...)
	}

	sort.Ints(sortedSli)

	fmt.Println("Result:", sortedSli)
}

func readInput(sli *[]int, sc *bufio.Scanner) {
	fmt.Print("Enter numbers separated by space: ")
	sc.Scan()
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Wrong input")
			os.Exit(0)
		}
	}()

	for _, i := range strings.Fields(sc.Text()) {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		} else {
			*sli = append(*sli, num)
		}
	}
}
