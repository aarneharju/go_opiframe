package main

import "fmt"

func main() {
	var b int = 15
	var a int
	var t int = 10

	// Initialization does not cover the whole array length. Rest are zero.
	numbers := [6]int{1,2,3,4}

	// For loops
	for a =0 ; a < 10; a++ {
		fmt.Printf("Value of  a:%d\n",a)
	}

	for(a < b) {
		a++
		fmt.Printf("Value of a:%d\n",a)
	}

	// i,x for index and value. Replace i with _ if you don't need index.
	for i,x := range numbers {
		fmt.Printf("Value of x=%d at %d\n",x,i)
	}

	// No parentheses around if condition. Braces are required (old languages didn't require parentheses for single line if statement. Which could cause heartbleed memory leaking bug for instance)
	if t > 5 {
		fmt.Printf("%d is bigger than 5\n",t);
	}

	if t < 5 {
		fmt.Printf("%d is smaller than 5\n",t);
	} else {
		fmt.Printf("%d is greater than 5\n",t);
	}

	// Switch has more uses than in most languages. You can compare all sorts of things.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	
	whatAmI := func(i interface{}) {
		switch r := i.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Printf("Don't recognize type %T\n",r)
		}
	}

	whatAmI(true)
	whatAmI(5)
	whatAmI("Hello")

}