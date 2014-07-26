package main

import "fmt"
import "net/http"
import "time"
import "strconv"

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

