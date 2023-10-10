package model

import (
	"encoding/csv"
	"strconv"
    "os"
	"time"
)

type ExampleData struct {
    Name string 
    Created time.Time 
    Done bool 
    Id int 
}
func (data ExampleData) Csv() string {
    return strconv.Itoa(data.Id) + "," + data.Name + "," + data.Created.Format("2006-01-02 15:04:05") + "," + strconv.FormatBool(data.Done)
}

// Reads the data from a csv file.
// todo: Maybe replace this with database access?
// update: Not used after sqlx impl
func Read(path string) []ExampleData {
    var data []ExampleData
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    // make sure file closes 
    defer file.Close()
    reader, err := csv.NewReader(file).ReadAll()
    if err != nil {
        panic(err)
    }
    for i, row := range reader {
        var name string
        var created time.Time
        var done bool
        id := i + 1
        for j, value := range row {
            switch j {
            case 0: name = value
            case 1: created, _ = time.Parse("2006-01-02 15:04:05", value)
            case 2: done, _ = strconv.ParseBool(value)
            } 
        }
        data = append(data, ExampleData{Name: name, Created: created, Done: done, Id: id})
    }
    return data
}

// Writes the data to a csv file.
// todo: Maybe replace this with database access?
// update: Not used after sqlx impl
func Save(data []ExampleData) {
    file, err := os.Create("example.csv")
    if err != nil {
        panic(err)
    }
    // make sure file closes 
    defer file.Close()
    writer := csv.NewWriter(file)
    for _, row := range data {
        writer.Write([]string{row.Name, row.Created.Format("2006-01-02 15:04:05"), strconv.FormatBool(row.Done)})
    }
    writer.Flush()
}
