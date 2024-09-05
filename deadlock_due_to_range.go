package main

import "fmt"

func deadlock_due_to_range() {
	goChannel := make(chan int) // unbuffered so cannot keep stuff
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}(goChannel)

	// now try to use a range based for loop to iterate over this channel
	// remember that range based loops keep going forever unless they are notified by the sender
	// that no more values are coming through channel closure from the sender!
	for value := range goChannel {
		fmt.Println(value)
	}
}
