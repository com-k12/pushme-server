package main

import "fmt"
import "net/http"
import "html"
import "log"

func handleGET (response_writer http.ResponseWriter, r *http.Request) {
    fmt.Fprintf (response_writer, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
    fmt.Printf ("Pushme server is running\n")

    http.HandleFunc("/", handleGET)

    log.Fatal (http.ListenAndServe("192.168.1.143:7000", nil))
}

