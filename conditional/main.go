package main 

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	x := 0
	
	
	if x == 4 {
		fmt.Println(4)
		
		} else if x == 5 {
			fmt.Println(5)
	} else if x == 6 {
		fmt.Println(6)
	} else {
		fmt.Println("other")
	}
}

func bar() {
	x := 4
	switch x {
		case 4: fmt.Println(4)
		case 5: fmt.Println(5)
		case 6: fmt.Println(6)
}
}