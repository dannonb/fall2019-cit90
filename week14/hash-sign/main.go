package main 

import (
	"crypto/sha256"
	"strings"
	"fmt"
	
)

const secret = "astronaut juggernaut jeigermeister"

func main() {
	email := "dannon@gmail.com"

	s := email + secret 
	sum := sha256.Sum256([]byte(s))
	hash := fmt.Sprintf("%x", sum)
	//valuetostoreincookie := email+"|"+hash

	 emailTamperedInCookie := "notdannon@gmail.com"
	 valueFromCookie := emailTamperedInCookie + "|" + hash

	xs := strings.Split(valueFromCookie, "|")
	emailFromCookie := xs[0]
	hashCodeFromCookie := xs[1]

	s2 := emailFromCookie + secret
	sum2 := sha256.Sum256([]byte(s2))
	hash2 := fmt.Sprintf("%x", sum2)

	if hashCodeFromCookie != hash2 {

	}
}