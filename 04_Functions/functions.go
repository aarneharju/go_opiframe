package main

import "fmt"

// Basic functions. Variable name before type in parameters and return value last
func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

// Multiple return values

func math_action(a int, b int, action string) (int, string) {
	if action == "add" {
		return a + b, "added"
	}

	if action == "subtract" {
		return a-b, "subtracted"
	}
	return 0, "unknown action"
}

// Variadic functions

func sum(vals ...int) int {
	total := 0
	for _, num := range vals { // index not needed, so using _
		total += num
	}
	return total
}

func main() {
	c := add(20, 10)
	d := subtract(20,10)

	fmt.Println("20+10 = ", c)
	fmt.Println("20-10 = ", d)
	
	e, act := math_action(20,10,"add")
	f, act2 := math_action(20,10,"subtract")
	_, act3 := math_action(20,10,"divide")

	fmt.Printf("20 %s to 10 = %d\n", act, e)
	fmt.Printf("10 %s from 20 = %d\n", act2, f)
	fmt.Printf("%s\n", act3)

	t := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Add together these numbers", t)
	fmt.Printf("Sum is %d\n", sum(t...))
}