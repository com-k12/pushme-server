package main

import "io/ioutil"
import "fmt"
import "bytes"
import "strconv"
import "strings"

// Глобальная конфигурация демона
var configuration = map[string]string {
    "server_address": "127.0.0.1",
    "server_port"   : "1904",
}

//------------------------------------------------------------------------------
// readConfiguration читает конфигурацию из файла
// TODO Выбрать формат конфигурации сервера (JSon, ini, hand-made)
func readConfiguration(file_name string) bool {
    file_content, err := ioutil.ReadFile (file_name)
    if err != nil {
        fmt.Println ("Error. Can't read configuration from file:", file_name)
        return false
    }

    // Читаем конфигурацию из файла.
    lines := bytes.Split (file_content, []byte{'\n'})

    for _, line := range lines {
        if (len(line) <= 1) {
            continue
        }
        if (line[0] == '#') {
            continue
        }

        name_value := bytes.Split (line, []byte{'='})

        if len(name_value) != 2 && len(name_value[0]) != 0 {
            fmt.Println ("Error. Wrong format of configuration file")
            return false
        }

        name := string(bytes.Trim(name_value[0], " "))

        if (len(name)>0) {
            value := string(bytes.Trim(name_value[1], " "))
            configuration[name] = value
        }
    }

    return true
}

//------------------------------------------------------------------------------
// getConfUint32 возвращает значение из конфигурации, если оно присутсвует
//               иначе возвращате default_value
func getConfUint32(name string, default_value uint32) uint32{
    temp_str, ok := configuration[name]

    if !ok {
        return default_value
    }

    temp_uint, err := strconv.ParseUint(strings.Trim(temp_str, " "), 10, 32)
    if err != nil {
        fmt.Printf ("Error. Incorect value of \"%s\"", name)
        return default_value
    }
    return uint32(temp_uint)
}
