package main

import (
	"io"
	"net/http"
	
)


func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func addCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name: name,
		Value: value,

	}
	http.SetCookie(w, &cookie)
}

func index(w http.ResponseWriter, req *http.Request) {
	addCookie(w, "name", "value")
	io.WriteString(w, "Welcome")
}