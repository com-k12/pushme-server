package main

import "fmt"
import "net/http"
import "log"
import "time"
import "container/list"
import "strconv"

const time_format = "15:04 02.01"

type Notification struct {
    date_time    time.Time
    level        int
    message      string
    readed_users map[string] bool
}

var notifications = list.New()
var notifications_index map[string] *Notification

//-----------------------------------------------------------------------------
// Обработка основного запроса
func handleGET (response_writer http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()

    buf, ok := values["user"]
    if (!ok) {
        fmt.Fprintf (response_writer, "User wasn't presented in request")
        return
    }
    user := buf[0]

    fmt.Fprintf (response_writer, "user: %s\n", user)
    //fmt.Fprintf (response_writer, "level: %s\n", level)
    //fmt.Fprintf (response_writer, "message: %s\n", message)
    //fmt.Fprintf (response_writer, "time: %s\n", time.Now().Format(time_format))
}

//-----------------------------------------------------------------------------
func main() {
    fmt.Printf ("Pushme server is running\n")

    http.HandleFunc("/", handleGET)
    http.HandleFunc("/set/", handleSET)

    log.Fatal (http.ListenAndServe("192.168.1.143:7000", nil))
}


//-----------------------------------------------------------------------------
// Обработка основного запроса
func handleSET (response_writer http.ResponseWriter, r *http.Request) {
    var notification Notification

    values := r.URL.Query()

    buf, ok := values["user"]
    if (!ok) {
        fmt.Fprintf (response_writer, "User wasn't presented in request")
        return
    }
    user := buf[0]

    buf, ok = values["level"]
    if (!ok) {
        fmt.Fprintf (response_writer, "Level wasn't presented in request")
        return
    }
    notification.level,_ = strconv.Atoi(buf[0])

    buf, ok = values["message"]
    if (!ok) {
        fmt.Fprintf (response_writer, "Message wasn't presented in request")
        return
    }
    notification.message = buf[0]
    notification.date_time = time.Now()

    fmt.Fprintf (response_writer, "URL: %s\n", r.URL.Path)
    fmt.Fprintf (response_writer, "user: %s\n", user)
    fmt.Fprintf (response_writer, "level: %i\n", notification.level)
    fmt.Fprintf (response_writer, "message: %s\n", notification.message)
    fmt.Fprintf (response_writer, "time: %s\n", notification.date_time.Format(time_format))

}
