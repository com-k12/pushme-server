package main

import "fmt"
import "net/http"

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
