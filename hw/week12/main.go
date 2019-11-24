package main

import (
	"io"
	"net/http"
	
)


func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func addCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name: name,
		Value: value,

	}
	http.SetCookie(w, &cookie)
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie")
	if err != nil {
		w.Write([]byte("error"))
	} else {
		w.Write([]byte(c.Value))
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	addCookie(w, "cookie name", "cookie value")
	
	io.WriteString(w, "Welcome")
	readCookie(w, req)
}