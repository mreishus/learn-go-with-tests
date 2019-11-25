// Package main provides ...
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	fmt.Println("Hello, world!")
	Greet(os.Stdout, "Person")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
