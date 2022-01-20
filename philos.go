/*
Implement the dining philosopher’s problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks,
with one chopstick between each adjacent pair of philosophers.

2. Each philosopher should eat only 3 times
(not in an infinite loop as we did in lecture)

3. The philosophers pick up the chopsticks in any order,
not lowest-numbered first (which we did in lecture).

4. In order to eat, a philosopher must get permission
from a host which executes in its own goroutine.

5. The host allows no more than 2 philosophers to eat concurrently.

6. Each philosopher is numbered, 1 through 5.

7. When a philosopher starts eating (after it has obtained necessary locks)
it prints “starting to eat <number>” on a line by itself,
where <number> is the number of the philosopher.

8. When a philosopher finishes eating (before it has released its locks)
it prints “finishing eating <number>” on a line by itself,
where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

var wg sync.WaitGroup
var host = make(chan bool, 2)

func (p Philo) eat() {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		//time.Sleep(1 * time.Second) //uncomment for demonstrating

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Printf("finishing eating %d\n", p.number)
		//time.Sleep(1 * time.Second) //uncomment for demonstrating

		<-host
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
