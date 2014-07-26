package main

import "fmt"
import "net/http"
import "log"
import "time"

const time_format = "15:04:05 02.01"

type Notification struct {
    date_time    time.Time
    level        int
    message      string
}

var notifications []Notification
var index = make(map[string] []Notification)

//-----------------------------------------------------------------------------
func main() {
    if !readConfiguration("pushme-server.conf") {
        return
    }

    fmt.Printf ("Pushme server is running\n")

    readUserConfiguration ()

    http.HandleFunc("/", handleGet)
    http.HandleFunc("/add/", handleAdd)
    http.HandleFunc("/register/", handleRegister)

    server_address, ok := configuration["server_address"]
    if (!ok) {
        log.Fatal ("No variable \"server_address\" in configuration")
    } else {
        log.Fatal (http.ListenAndServe(server_address, nil))
    }
}

