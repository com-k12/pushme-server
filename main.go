package main

import "fmt"
import "net/http"
import "log"
import "time"

const layout = "15:04 02.01"

func handleGET (response_writer http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()
    fmt.Fprintf (response_writer, "URL: %s\n", r.URL.Path)
    fmt.Fprintf (response_writer, "user: %s\n", values["user"][0])
    fmt.Fprintf (response_writer, "message: %s\n", values["message"][0])
    fmt.Fprintf (response_writer, "time: %s\n", time.Now().Format(layout))
}

func main() {
    fmt.Printf ("Pushme server is running\n")

    http.HandleFunc("/", handleGET)

    log.Fatal (http.ListenAndServe("192.168.1.143:7000", nil))
}

