/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should
take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := RequestAccelaration(reader)
	o_v := RequestInitialVelocity(reader)
	o_s := RequestInitialDisplacement(reader)
	fn := GenDisplaceFn(a, o_v, o_s)
	t := RequestTime(reader)
	fmt.Printf("Displacement after %f seconds: %f\n", t, fn(t))
}

func GenDisplaceFn(a, o_v, o_s float64) func(float64) float64 {
	s := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + o_v*t + o_s
	}
	return s
}

func RequestAccelaration(r *bufio.Reader) float64 {
	var inputA float64
	for {
		fmt.Print("Enter accelaration (type \"X\" - to exit): ")
		inputString, _ := r.ReadString('\n')
		if strings.TrimSpace(inputString) == "X" {
			fmt.Println("Exiting a program.")
			os.Exit(0)
		} else {
			tmp, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
			if err != nil {
				fmt.Println("Wrong input!")
				continue
			} else {
				inputA = tmp
				break
			}
		}
	}
	return inputA
}

func RequestInitialVelocity(r *bufio.Reader) float64 {
	var inputV0 float64
	for {
		fmt.Print("Enter initial velocity (type \"X\" - to exit): ")
		inputString, _ := r.ReadString('\n')
		if strings.TrimSpace(inputString) == "X" {
			fmt.Println("Exiting a program.")
			os.Exit(0)
		} else {
			tmp, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
			if err != nil {
				fmt.Println("Wrong input!")
				continue
			} else {
				inputV0 = tmp
				break
			}
		}
	}
	return inputV0
}

func RequestInitialDisplacement(r *bufio.Reader) float64 {
	var inputS0 float64
	for {
		fmt.Print("Enter initial displacement (type \"X\" - to exit): ")
		inputString, _ := r.ReadString('\n')
		if strings.TrimSpace(inputString) == "X" {
			fmt.Println("Exiting a program.")
			os.Exit(0)
		} else {
			tmp, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
			if err != nil {
				fmt.Println("Wrong input!")
				continue
			} else {
				inputS0 = tmp
				break
			}
		}
	}
	return inputS0
}

func RequestTime(r *bufio.Reader) float64 {
	var inputT float64
	for {
		fmt.Print("Enter time (type \"X\" - to exit): ")
		inputString, _ := r.ReadString('\n')
		if strings.TrimSpace(inputString) == "X" {
			fmt.Println("Exiting a program.")
			os.Exit(0)
		} else {
			tmp, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
			if err != nil {
				fmt.Println("Wrong input!")
				continue
			} else {
				inputT = tmp
				break
			}
		}
	}
	return inputT
}
