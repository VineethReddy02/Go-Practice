package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id                     int
	lChopStick, rChopStick *ChopStick
}

func main() {
	sticks := initSticks()
	philosophers := initPhilosophers(sticks)
	wantsToEat := make(chan int)
	finishedEating := make(chan int)

	go host(wantsToEat, finishedEating)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(wantsToEat, finishedEating, &wg)
	}
	wg.Wait()
}

// It makes a little sense to have a "host" for two dining philosophers.
// if we want two dining guys anyway, we should introduce
// two hosts (instead of this suboptimal switching)

func host(hungryPhilosophers chan int, fedPhilosophers chan int) {
	var a, b sync.Mutex

	counter := 0
	for {
		<-hungryPhilosophers
		a.Lock()
		counter++
		if counter < 15 {
			b.Lock()
			<-hungryPhilosophers
			counter++
			<-fedPhilosophers
			b.Unlock()
		}
		<-fedPhilosophers
		a.Unlock()
	}
}

func initPhilosophers(sticks []*ChopStick) []*Philosopher {
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = new(Philosopher)
		philosophers[i].id = i + 1
		philosophers[i].lChopStick = sticks[i]
		philosophers[i].rChopStick = sticks[(i+1)%5]
	}
	return philosophers
}

func initSticks() []*ChopStick {
	sticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopStick)
	}
	return sticks
}

func (phil Philosopher) eat(wantsToEat chan int, finishedEating chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		wantsToEat <- phil.id
		phil.lChopStick.Lock()
		phil.rChopStick.Lock()
		fmt.Println("Starting to eat ", phil.id)
		fmt.Println("Finishing eating ", phil.id)
		phil.rChopStick.Unlock()
		phil.lChopStick.Unlock()
		finishedEating <- phil.id
	}
}
