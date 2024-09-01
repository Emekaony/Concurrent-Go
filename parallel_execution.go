package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func valueSet3(goChannel chan<- int) {
	defer wg.Done()
	for i := 10; i < 15; i++ {
		goChannel <- i
	}
}

func valueSet2(goChannel chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		goChannel <- i
	}
}

func valueSet1(goChannel chan<- int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		goChannel <- i
	}
}

func main2() {
	runtime.GOMAXPROCS(4)
	wg.Add(3)
	goChannel := make(chan int, 20)
	go valueSet1(goChannel)
	go valueSet3(goChannel)
	go valueSet2(goChannel)

	// this works BUT there's an even better way using range based loops to avoid unused values
	// for i := 0; i < 15; i++ {
	// 	fmt.Println(<-goChannel)
	// }

	// good practice to use range pased loop to avoid orphaned values sitting unused in the channel buffer.
	for val := range goChannel {
		fmt.Println(val)
	}

	wg.Wait()
	close(goChannel) // it is good practice to close channels when ur done with them
}
