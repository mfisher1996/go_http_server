package main

import (
	"bytes"
	"html/template"
	"http_test/model"
	"io"
	"net/http"
	"time"
)

var tpl *template.Template
var err error

func init() {
    tpl = template.Must( template.ParseGlob("view/*.html"))
}

func hello(w http.ResponseWriter, r *http.Request) {
    println("Call to hello")
    exampleData := []model.ExampleData{{ Name: "Mow Lawn", Created: time.Now() , Done: true, Id: 1}, { Name: "Take Out Trash", Created: time.Now(), Done: false, Id: 2}}
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    err := tpl.ExecuteTemplate(writer,"view.html", exampleData)
    if err != nil {
        println("Error with template ", err.Error())
    }
    w.Write(buf.Bytes())
}

func main() {
    http.Handle("/view/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/hello", hello)
    //http.HandleFunc("/save", save)
    http.HandleFunc("/edit", edit)
    http.ListenAndServe(":8080", nil)
}

//func save(w http.ResponseWriter, r *http.Request) {
    //var buf bytes.Buffer
    ////writer := io.Writer(&buf)
    //w.Header().Set("Content-Type", "text/html")
    //tpl.Parse(`{{ template "row_item" . }}`)
    ////tpl.Execute(writer)
    //w.Write(buf.Bytes())
//}

func edit(w http.ResponseWriter, r *http.Request) {
    println("Call to edit")
    err := r.ParseForm()
    if err != nil {
        panic(err)
    }
    created, _ := time.Parse("", r.PostForm.Get("Created"))
    done := r.PostForm.Get("Done") == "true"
    exampleData := model.ExampleData{Name: r.PostForm.Get("Name"), Created: created , Done: done}
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    tpl.Parse(` {{ template "row_item_as_form". }} `)
    tpl.Execute(writer, exampleData)
    w.Write(buf.Bytes())
}
