/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.

The race condition is the problem when your program and its outcome
depends on interleavings (e. g. multiple goroutines are trying to access and manipulate the same variable).
And since the interleavings are non-deterministic your output result will be non-deterministic too.
*/

package main

import (
	"fmt"
	"time"
)

var x int

func main() {

	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine #1")
		x = x + 1
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine #2")
		x = x + 1
	}()

	time.Sleep(3 * time.Second)
}
