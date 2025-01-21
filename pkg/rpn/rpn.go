package rpn

import (
    "net/http"
    "os"
    "html/template"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

func Main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    // Добавьте следующие две строки
    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}