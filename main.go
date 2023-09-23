package main

import (
	"bytes"
	"html/template"
	"http_test/model"
	"io"
	"net/http"
	"time"
)

const TEMPLATE = `
    {{define "item"}}
        <tr><td>{{.Name}}</td><td>{{.Created.Format "2006-01-02 15:04:05" }}</td><td><input type="checkbox" value="{{.Done}}"><label>Done</label></input></td>
    {{end}}
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        <h1>Hello World</h1>
        <table>
    {{range .}}
    <tr>
        {{template "item" .}}
    </tr>
    {{end}}

        </table>
    </body>
    `

func hello(w http.ResponseWriter, r *http.Request) {
    exampleData := []model.ExampleData{{ Name: "Mow Lawn", Created: time.Now() }, { Name: "Take Out Trash", Created: time.Now() }}
    
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")

    //io.WriteString(writer,`
//<html>
    //<head>
        //<title>Hello World</title>
    //</head>
    //<body>
        //<h1>Hello World</h1>
        //<table>
        //`)
    //model.AllAsHtml( &exampleData, &writer)
    //io.WriteString(writer,`
        //</table>
    //</body>`)

    t := template.Must(template.New("html").Parse(TEMPLATE))
    t.Execute(writer, exampleData)
    w.Write(buf.Bytes())
    
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)

    // test stuff
}
