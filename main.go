package main

import "fmt"
import "net/http"
import "log"
import "time"
import "strconv"

const time_format = "15:04 02.01"

type Notification struct {
    date_time    time.Time
    level        int
    message      string
    readed_users map[string] bool
}

var notifications []Notification
var index = make(map[string] []Notification)

//-----------------------------------------------------------------------------
// Обработка основного запроса
func handleGet (response_writer http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()

    buf, ok := values["user"]
    if (!ok) {
        fmt.Fprintf (response_writer, "User wasn't presented in request")
        return
    }
    user := buf[0]

    buf_list, ok := index[user]
    if (!ok) {
        fmt.Fprintf (response_writer, "User %s not registered", user)
        return
    }

    for i:= range buf_list {
        notification := buf_list[i]
        fmt.Fprintf (response_writer, "%s|%d|%s", notification.date_time.Format(time_format), notification.level , notification.message)
    }
    buf_list = make([]Notification,0)
    //fmt.Fprintf (response_writer, "level: %s\n", level)
    //fmt.Fprintf (response_writer, "message: %s\n", message)
    //fmt.Fprintf (response_writer, "time: %s\n", time.Now().Format(time_format))
}

//-----------------------------------------------------------------------------
func main() {
    fmt.Printf ("Pushme server is running\n")

    http.HandleFunc("/", handleGet)
    http.HandleFunc("/add/", handleAdd)

    log.Fatal (http.ListenAndServe("192.168.1.143:7000", nil))
}


//-----------------------------------------------------------------------------
// Обработка запроса на добавление нотификации
func handleAdd (response_writer http.ResponseWriter, r *http.Request) {
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
    fmt.Fprintf (response_writer, "level: %d\n", notification.level)
    fmt.Fprintf (response_writer, "message: %s\n", notification.message)
    fmt.Fprintf (response_writer, "time: %s\n", notification.date_time.Format(time_format))

    notifications = append (notifications, notification)
    _,ok = index[user]
    if (!ok) {
        fmt.Println("New user")
        index[user] = make([]Notification,0)
        index[user] = append(index[user], notification)
    } else {
        fmt.Println("Yes user")
        index[user] = append(index[user], notification)
    }
}
