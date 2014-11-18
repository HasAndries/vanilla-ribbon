package main

import (
	"net/http"
	"os"
	"path"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]string)
	if params["page"] = r.URL.Path[1:]; params["page"] == "" {
		params["page"] = "index"
	}
	layoutPath := path.Join("views", "layout.html")
	contentPath := path.Join("views", params["page"]+".html")

	if _, err := os.Stat(contentPath); err != nil {
		contentPath = path.Join("views", "not-found.html")
	}

	tmpl, err := template.ParseFiles(layoutPath, contentPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
