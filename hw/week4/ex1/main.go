package main 

import (
	"fmt"
)

type person struct{
	first string
	last string
	favs []string
}

func main() {
	p1 := person{
		"dannon",
		"bigham",
		[]string{"choco", "vanilla"},
	}

	p2 := person{
		"destiny",
		"malone",
		[]string{"birthday cake", "berry"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favs {
		fmt.Println(i, v)
	}

	for i, v := range p2.favs {
		fmt.Println(i, v)
	}
}