package main 

import (
	"os"
	"fmt"
)

func main() {
	file := "filename"
	file2 := "newfilename"
	os.Rename(file, file2)
	fmt.Printf(file)
}