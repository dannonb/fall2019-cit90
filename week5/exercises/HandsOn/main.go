package main

import "fmt"

func main() {
	// exercise 6
	func ()  {
		fmt.Println("")
		fmt.Println("EXERCISE 6")
		fmt.Println("dan")
		fmt.Println("")
	}()
	// exercise 7
	a := func() int {
		return 6
	}
	fmt.Println("EXERCISE 7")
	fmt.Println(a())
	fmt.Println("")
	
	// exercise 8
	s := bar()
	fmt.Println("EXERCISE 8")
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println("")

	// exercise 9
	r := func(x int) int{
		return x
		}
	z := dan(r, 7)
	fmt.Println("EXERCISE 9")
	fmt.Println(z)
	fmt.Println("")
	
	// exercise 10
	p := bar()
	fmt.Println("EXERCISE 10")
	fmt.Println(p())
	fmt.Println("")

	
	// pointer exercises
	fmt.Println("POINTER EXERCISES")
	x := 5
	fmt.Println(&x)
	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",&x)
	y := &x
	fmt.Println(y)
	fmt.Printf("%T\n",y)
	*y = 6
	fmt.Println(x)
}

// exercises 8 & 10
func bar() func() int{
	x := 8
	return func() int{
		x++
		return x
	}
}

// exercise 9

func dan(f func (x int) int, y int) int{
r := f(y)
return r
}







