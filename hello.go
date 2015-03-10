package main

import (
    "net/http" //package for http based web programs
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Handling HTTP request")
    fmt.Fprintf(w, "Hello world!")
}

func main() {
    fmt.Println("Initializing Hello World...")
    http.HandleFunc("/", handler) // redirect all urls to the handler function
    fmt.Println("DONE setting handler...")
    http.ListenAndServe(":8080", nil) // listen for connections at port 9999 on the local machine
    fmt.Println("DONE Initializing Hello World...")
}
