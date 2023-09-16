package main

import (
    "net/http" 
    "io"
    "http_test/model"
    "bytes"
)

func hello(w http.ResponseWriter, r *http.Request) {
    exampleData := []model.ExampleData{{ Name: "John", Age: 30 }, { Name: "Jane", Age: 25 }}
    
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    io.WriteString(writer,`
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        <h1>Hello World</h1>
        <table>
        `)
    model.AllAsHtml( &exampleData, &writer)
    io.WriteString(writer,`
        </table>
    </body>`)
    w.Write(buf.Bytes())
    
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)

    // test stuff
}
