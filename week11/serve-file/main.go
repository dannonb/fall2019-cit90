package main 

import (
	"net/http"
	"io"
	// "os"

)

func main() {
	// NOTE: handle func is a configuration funciton meaning that it must be called before starting the server 
	http.Handle("/file/", http.FileServer(http.Dir("./assets")))
	http.HandleFunc("/", index)
	//http.HandleFunc("/dog.jpg", toby)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<img src="/file/img/dog.jpg" alt="doggo">
	</body>
	</html>`)
}

// func toby(w http.ResponseWriter, r *http.Request) {
// 	f, err := os.Open("toby.jpeg")

// 	if err != nil {
// 		http.Error(w, "coulnt open file", http.StatusInternalServerError)
// 	}
// 	defer f.Close()

// 	io.Copy(w, f)
// }


// func toby(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "dog.jpg")
// }