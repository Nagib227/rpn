package rpn

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("C:\\Users\\leoni\\Documents\\Practice\\GO\\my-go-project\\WebApp\\rpn\\cmd\\index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
func indexHandlerCreation(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("assets\\html\\creation.html")).Execute(w, nil)
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
	mux.HandleFunc("/creation", indexHandlerCreation)
	http.ListenAndServe(":"+port, mux)
}
