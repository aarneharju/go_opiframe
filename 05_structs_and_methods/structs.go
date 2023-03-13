package main

import "fmt"

// Create a couple of struct types

type dog struct {
	name, breed string
}

type cat struct {
	name, breed string
}

type animal interface {
	move()
	speak()
}

func (d dog) move() {
	fmt.Printf("%s runs. See %s run\n", d.name, d.name)
}

func (d dog) speak() {
	fmt.Printf("%s barks. %s is a %s\n", d.name, d.name, d.breed)
}

func (c cat) move() {
	fmt.Printf("%s sneaks. See %s sneak\n", c.name, c.name)
}

func (c cat) speak() {
	fmt.Printf("%s meows. %s is a %s\n", c.name, c.name, c.breed)
}

func act(a animal) {
	a.move()
	a.speak()
}

func main() {
	duke := dog{name:"Duke", breed:"Sitter"}
	whiskers := cat{name:"Whiskers", breed:"Persian"}

	fmt.Println("Accessing through struct type instance")
	
	duke.move()
	duke.speak()
	whiskers.move()
	whiskers.speak()
	
	fmt.Println("Accessing through interface")

	act(duke) // structural typing makes this work
	act(whiskers)

}