package main

import (
    "net/http" //package for http based web programs
    "fmt"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Handling HTTP request")
    fmt.Fprintf(w, "Hello world!")
}

func main() {
    fmt.Println("Initializing Hello World...")
    http.HandleFunc("/", handler) // redirect all urls to the handler function
    fmt.Println("DONE setting handler...")
    port := "8080" // default to port 8080
    if len(os.Args) > 1 {
        port = os.Args[1] // command line argument can be provided for port
    }
    http.ListenAndServe(":" + port, nil) // listen for connections at port on the local machine
    fmt.Println("DONE Initializing Hello World...")
}
