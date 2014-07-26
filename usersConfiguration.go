package main

import "fmt"
import "io/ioutil"
import "bytes"

func readUserConfiguration() {
    users_file_name, ok := configuration["users_file_name"]
    if (!ok) {
        fmt.Println ("No variable \"users_file_name\" in configuration")
    }

    file_content, _ := ioutil.ReadFile (users_file_name)

    // Читаем конфигурацию из файла.
    lines := bytes.Split (file_content, []byte{'\n'})

    for _, line := range lines {
        if (len(line) <= 1) {
            continue
        }
        if (line[0] == '#') {
            continue
        }

        if (len(line)>0) {
            index[string(line)] = make([]Notification, 0)
        }
    }
}
