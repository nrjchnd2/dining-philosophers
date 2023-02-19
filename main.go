package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Obama", leftFork: 1, rightFork: 0},
	{name: "Plato", leftFork: 2, rightFork: 1},
	{name: "Socrates", leftFork: 3, rightFork: 2},
	{name: "Aristotle", leftFork: 4, rightFork: 3},
	{name: "Picaso", leftFork: 0, rightFork: 4},
}

var hunger = 3 // number of times a philisopher needs to eat to finish his food
var eatingTime = 1 * time.Second
var thinkingTime = 3 * time.Second
var orderFinishedMutex sync.Mutex
var orderDiningFinished []string

func main() {
	fmt.Println("The Dining Philosophers probelem")
	fmt.Println("--------------------------------")
	fmt.Println("\n The table is empty")

	dine()

	fmt.Println("The table is empty")

	fmt.Printf("Order fisnihed in %+v\n", orderDiningFinished)

}

func dine() {
	//waitgroup to wait for all philosophers to finish eating
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	//waitgroup to wait for all philosophers to sit at the table before startig to eat
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, seated, forks)
	}
	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, seated *sync.WaitGroup, forks map[int]*sync.Mutex) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table \n", philosopher.name)
	//this philosopher is seated
	seated.Done()
	//wait for all philosophers to be seated
	seated.Wait()

	for i := hunger; i > 0; i-- {
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s has grabbed the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s has grabbed the left fork\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s has grabbed the left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s has grabbed the right fork\n", philosopher.name)
		}
		fmt.Printf("\t%s has both the forks and has now started eating\n", philosopher.name)
		time.Sleep(eatingTime)

		fmt.Printf("\t%s is thinking\n", philosopher.name)
		time.Sleep(thinkingTime)

		//unlocking mutexes for both forks
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
	}
	fmt.Printf("%s is satisfied and has left the table\n", philosopher.name)

	orderFinishedMutex.Lock()
	orderDiningFinished = append(orderDiningFinished, philosopher.name)
	orderFinishedMutex.Unlock()
}
