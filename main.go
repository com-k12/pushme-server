package main

import "fmt"
import "net/http"
import "log"
import "time"
import "strconv"

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

    index["aleus"] = make([]Notification,0)
    index["nezloi"] = make([]Notification,0)
    index["sdimanx"] = make([]Notification,0)

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


//-----------------------------------------------------------------------------
// Обработка запроса на добавление нотификации
func handleAdd (response_writer http.ResponseWriter, r *http.Request) {
    var notification Notification

    values := r.URL.Query()

    buf, ok := values["user"]
    //if (!ok) {
    //    fmt.Fprintf (response_writer, "User wasn't presented in request")
    //    return
    //}
    //admin_user := buf[0]

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
    //fmt.Fprintf (response_writer, "admin_user: %s\n", admin_user)
    fmt.Fprintf (response_writer, "level: %d\n", notification.level)
    fmt.Fprintf (response_writer, "message: %s\n", notification.message)
    fmt.Fprintf (response_writer, "time: %s\n", notification.date_time.Format(time_format))

    notifications = append (notifications, notification)
    for i:= range index {
        index[i] = append(index[i], notification)
    }
}

//------------------------------------------------------------------------------
// Регистрирует пользователя в системе
func handleRegister (response_writer http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()

    buf, ok := values["user"]
    if (!ok) {
        fmt.Fprintf (response_writer, "User wasn't presented in request")
        return
    }
    user := buf[0]

    fmt.Println("New user", user)
    index[user] = make([]Notification,0)
}

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
        fmt.Fprintf (response_writer, "%s|%d|%s\n", notification.date_time.Format(time_format), notification.level , notification.message)
    }
    index[user] = make([]Notification,0)
}
