package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

// Load the index.html template.
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))
func main() {
    // Serve / with the index.html file.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("test_ajax.html")
        t.Execute(w, "gan")
    })

    // Serve /callme with a text response.
    http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
        gan:=[]string{"1","2","3"}
        fmt.Fprintln(w, gan)
    })

    // Start the server at http://localhost:9000
    log.Fatal(http.ListenAndServe(":8000", nil))
}
