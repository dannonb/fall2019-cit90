package main 

import (
	"fmt"
	"./helpers"
)

func main() {
	fmt.Println("hey")
	x := Foo()
	fmt.Println(x)
	y := helpers.Bar()
	fmt.Println(y)
}