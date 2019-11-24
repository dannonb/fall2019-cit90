package main 

import "fmt"

type hotdog int 
type person struct {
	name string
	age int
}

type sa struct {
	person 
	ltk bool
}



var dan hotdog



func main () {
	dan := 4
	
	d := sa{
	person: person{
		name: "dannon",
		age: 25,
	},
	ltk: false,

	}
	fmt.Println(d)
		d.Yeah("bigham")
	}

func (a sa) Yeah(n string) {
		fmt.Println("dannon",n)
	}

func (h hotdog) Eat(s sa) {
	fmt.Printf("")
}
