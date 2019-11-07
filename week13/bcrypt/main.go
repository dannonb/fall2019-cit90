package main 
import (
	"fmt"
	""
)

func main() {
p := "bananabread"

bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
if err != nil {
	log.Fatalln("error")
}
fmt.Println(bs)

p2 := "bananabread25"
err := bcrypt.CompareHashAndPassword(bs, []byte(p2))
if err != nil {
	log.Fatalln("passwords dont match")
}
fmt.Println("passwords match")
}