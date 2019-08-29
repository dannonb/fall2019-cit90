
// multiple ways to run a for loop
package main

import "fmt"

func main() {
Dan()
}

func foo() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 10000 {
			break
		}
	}
}

func Bar() {
	i := 0
	for i < 100 {
		fmt.Println(i)
		i++
	}
}

func Dan() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

// for init; cond; post {}