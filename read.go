/*
Write a program which reads information from a file and
represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and
after all lines have been read from the file,
your program will have a slice containing one struct for
each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs
and print the first and last names found in each struct.
*/

/*the text file that was used for testing:

Phoenix Armitage
Jamie Dorsey
Kaitlin Ferguson
Asiya Millington
Davey Bass
Celeste Burt
Ryker Foley
Lily-Grace Potter
Tj Lowe
Tyler Monaghan

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}

	var fnameSlice, lnameSlice []string
	var resultSlice []Name

	fmt.Println("Enter file path: ")
	pathReader := bufio.NewReader(os.Stdin)
	inputPath, _ := pathReader.ReadString('\n')
	inputPath = strings.TrimSpace(inputPath)

	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanWords)
	position := 1

	for fileScanner.Scan() {
		if position%2 != 0 {
			fnameSlice = append(fnameSlice, fileScanner.Text())
		} else {
			lnameSlice = append(lnameSlice, fileScanner.Text())
		}
		position++
	}

	for i := 0; i < len(fnameSlice); i++ {
		resultSlice = append(resultSlice, Name{fnameSlice[i], lnameSlice[i]})
	}

	fmt.Printf("Result slice of structs: %v\n", resultSlice)

	for i := range resultSlice {
		fmt.Printf("Struct %d: %s %s\n", i, resultSlice[i].fname, resultSlice[i].lname)
	}
}
