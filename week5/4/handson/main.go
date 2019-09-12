package main 

import "fmt"

func main() {
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