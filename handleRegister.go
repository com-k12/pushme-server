package main

import "fmt"
import "net/http"

//------------------------------------------------------------------------------
// Регистрирует пользователя в системе
func handleRegister (response_writer http.ResponseWriter, r *http.Request) {
    values := r.URL.Query()

    buf, ok := values["user"]
    if (!ok) {
        fmt.Fprintf (response_writer, "#error|05|User wasn't presented in request")
        return
    }
    user := buf[0]

    fmt.Println("New user", user)
    index[user] = make([]Notification, 0)
}

