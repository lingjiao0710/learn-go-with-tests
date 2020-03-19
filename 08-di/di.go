package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(buffer io.Writer, str string) {
	fmt.Fprintf(buffer, "Hello, %s", str)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Elodie")

	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
