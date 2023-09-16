package model

import (
    "html/template"; "io"
)

const ITEM_TEMPLATE = ``//Template
const LIST_TEMPLATE = `<table>
    <tr><th>Name</th><th>Age</th></tr>
    {{range . }}
    <tr><td>{{.Name}}</td><td>{{.Age}}</td></tr>
    {{end}}
</table>`//Template

type ExampleData struct {
    Name string
    Age int
}

func AllAsHtml(d *[]ExampleData, writer *io.Writer) {
    t := template.Must(template.New("html").Parse(LIST_TEMPLATE))
    t.ExecuteTemplate(*writer,"html", d)
}

func (e ExampleData) AsHtml(writer *io.Writer) {
    t := template.Must(template.New("html").Parse(LIST_TEMPLATE ))
    t.Execute(*writer, e)
}
