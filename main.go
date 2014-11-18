package main

import (
	"net/http"
	"path"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var pageName string
	if pageName = r.URL.Path[1:]; pageName == "" {
		pageName = "index"
	}
	lp := path.Join("views", "layout.html")
	fp := path.Join("views", pageName+".html")

	// Note that the layout file must be the first parameter in ParseFiles
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
