package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstname string
	lastName  string
	age       int
	weight    int
}

// create a new person!!!
func NewPerson(firstName, lastName string, age, weight int) *Person {
	return &Person{
		firstname: firstName,
		lastName:  lastName,
		age:       age,
		weight:    weight,
	}
}

func (p *Person) sayHello() {
	message := fmt.Sprintf("Hello there %s, welcome to the team", p.firstname)
	fmt.Println(message)
}

// OH MY GOD, golang is so wonderful!!!!
func main() {
	duration := 5 * time.Second
	time.Sleep(duration)
	fmt.Printf("I slept for %v and now I am alive\n", duration)
}
