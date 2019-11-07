package main 

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s1 := "this is text"
	s2 := "this is also text"

sum1 := sha256.Sum256([]byte(s1))
sum2 := sha256.Sum256([]byte(s2))

fmt.Println(sum1)
fmt.Println(sum2)
fmt.Println(sum1 == sum2)
}

/**/