package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page Not Found!", http.StatusFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported!", http.StatusFound)
	}
	fmt.Fprintf(w, "HELLO!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm Erro : %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successful!\n")
	name := r.FormValue("name")
	add := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", add)
}

func main() {
	fileServer := http.FileServer(http.Dir("PATH"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Server Started at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
