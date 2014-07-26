package main

import "fmt"
import "net/http"
import "time"
import "strconv"
import "strings"

//-----------------------------------------------------------------------------
// Обработка запроса на добавление нотификации
func handleAdd (response_writer http.ResponseWriter, r *http.Request) {
    var notification Notification

    values := r.URL.Query()

    buf, ok := values["level"]
    if (!ok) {
        fmt.Fprintf (response_writer, "#error|03|Level wasn't presented in request")
        return
    }
    notification.level,_ = strconv.Atoi(buf[0])

    buf, ok = values["message"]
    if (!ok) {
        fmt.Fprintf (response_writer, "#error|04|Message wasn't presented in request")
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
    buf, ok = values["users"]
    if (ok) {
        users := strings.Split(buf[0], ",")
        for i:=range users {
            _,ok = index[users[i]]
            if (ok) {
                index[users[i]] = append(index[users[i]], notification)
            }
        }
    } else {
        for i:=range index {
            index[i] = append(index[i], notification)
        }
    }
}

