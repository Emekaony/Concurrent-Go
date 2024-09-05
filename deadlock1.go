package main

func deadlock1() {
	goChan := make(chan int)
	goChan <- 1
}
