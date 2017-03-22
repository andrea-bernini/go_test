package main

import "net/http"
import "log"

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
    log.Printf("richiesta")	
}
