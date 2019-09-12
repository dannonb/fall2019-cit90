package main 

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	fmt.Println("x is stored at", &x)
	fmt.Printf("%T", x)

	var a *int
	a = &x 
	fmt.Println(a)

	
	fmt.Println("before foo", x)
	foo(&x)
	fmt.Println(x)


}

func foo(n *int) {
*n++
fmt.Println(n)
}

