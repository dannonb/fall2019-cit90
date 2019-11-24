package main 

import "fmt"

type person struct {
	first, last string
}

type storage interface {
	access(int) person
	set(int, person)
}

type dbPG map[int]person

type dbMongo map[int]person

func (db dbPG) access(x int) person{
	return db[x]
}
 
func (db dbPG) set(x int, p person) {
	fmt.Println()
}


func storer(s storage, x int, p person) {
	fmt.Println("storing a person in a db")
}

func main() {
dan := person{
	first: "dan",
	last: "big",
}

fmt.Println(dan)


}