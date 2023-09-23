package model

import (
    "html/template";
    "io";
    "time";
)

const ITEM_TEMPLATE = `<td>{{.Name}}</td><td>{{.Created.Format "2006-01-02 15:04:05" }}</td><td><checkbox value="{{.Done}}"></checkbox></td>`
const LIST_TEMPLATE = `<table>
    <tr><th>Name</th><th>Created</th></tr>
    {{range . }}
    <tr><td>{{.Name}}</td><td>{{.Created.Format "2006-01-02 15:04:05" }}</td></tr>
    {{end}}
</table>`//Template

type ExampleData struct {
    // todo task
    Name string
    Created time.Time
    Done bool

}

func AllAsHtml(d *[]ExampleData, writer *io.Writer) {
    t := template.Must(template.New("html").Parse(LIST_TEMPLATE))
    t.ExecuteTemplate(*writer,"html", d)
}

func (e ExampleData) AsHtml(writer *io.Writer) {
    t := template.Must(template.New("html").Parse(ITEM_TEMPLATE)) 
    t.Execute(*writer, e)
}
