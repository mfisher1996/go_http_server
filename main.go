package main

import (
    "net/http"; "io"
)
func hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w,`
            <html>
                <head>
                    <title>Hello World</title>
                </head>
                <body>
                    <h1>Hello World</h1>
                </body>
            </html>
            `)
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}
