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

	m := map[int]person{
		1: p1,
		2: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
	}
}