package main

import (
	"bytes"
	"html/template"
	"http_test/model"
	"http_test/sql"
	"io"
	"net/http"
	"strconv"
	"time"

	"slices"

	"github.com/jmoiron/sqlx"
)

var tpl *template.Template
var db sqlx.DB
var exampleData []model.ExampleData

func init() {
    db = sql.Init()
    tpl = template.Must( template.ParseGlob("view/*.html"))
}

func main() {
    http.Handle("/view/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/save", save)
    http.Handle("/edit/", http.StripPrefix("/edit/", http.HandlerFunc(edit)))
    http.HandleFunc("/new", new)
    http.ListenAndServe(":8080", nil)
}

func new(w http.ResponseWriter, r *http.Request) {
    println("Call to new")
    var ids []int
    err := db.Select(&ids, "SELECT MAX(id) FROM tasks")
    max_id:= slices.Max(ids)
    if err != nil {
        println("Error with select ", err.Error())
    }
    println("Max id is ", max_id)
    data := model.ExampleData{Id: max_id + 1, Created: time.Now(), Done: false}
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    tpl.ExecuteTemplate(writer, "new.html", data)
    w.Write(buf.Bytes())
}

func hello(w http.ResponseWriter, r *http.Request) {
    println("Call to hello")
    var exampleData =[]model.ExampleData{} 
    db.Select(&exampleData, "SELECT * FROM tasks")
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    err := tpl.ExecuteTemplate(writer,"view.html", exampleData)
    if err != nil {
        println("Error with template ", err.Error())
    }
    w.Write(buf.Bytes())
}

// This is the hanlder function for generating the editable row that will be 
// used to replace the row making the call. All data should be set as values. 
// The template used is 'edit.html'
func edit(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path
    println("Call to edit with id ", id)
    var datas = []model.ExampleData{}
    err := db.Select(&datas, "SELECT * FROM tasks WHERE id = $1;", id)
    if err != nil {
        println("Error with select ", err.Error())
    }
    data := datas[0]
    var buf bytes.Buffer
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    tpl.ExecuteTemplate(writer, "row-form.html", data)
    w.Write(buf.Bytes())
}

func save(w http.ResponseWriter, r *http.Request) {
    println("Call to save")
    var buf bytes.Buffer
    r.ParseForm()
    for key, value := range r.Form {
        println(key, value[0])
    }
    i, err := strconv.Atoi(r.Form.Get("Id"))
    if err != nil {
        println("Error with conversion ", err.Error())
        i = 0
    }
    data := model.ExampleData{Id: i, Name: r.Form.Get("Name"), Created: time.Now(), Done: r.Form.Get("Done") == "on"}
    _,err = db.Exec("INSERT INTO tasks (id, name, created, done) VALUES ($1, $2, $3, $4)", data.Id, data.Name, data.Created.Format("2006-01-02 15:04:05"), data.Done)
    if err != nil {
        println("Error with insert ", err.Error())
    }
    writer := io.Writer(&buf)
    w.Header().Set("Content-Type", "text/html")
    tpl.ExecuteTemplate(writer, "row.html", data)
    w.Write(buf.Bytes())
}

