package main

import (
	"net/http"
	"os"
	"path"
	"text/template"
	"io"
	"time"
)

const STATIC_URL string = "/public/"
const STATIC_ROOT string = "content/public/"
const VIEWS_ROOT string = "content"

func processStatic(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func processView(w http.ResponseWriter, req *http.Request) {
	params := make(map[string]string)
	if params["page"] = req.URL.Path[1:]; params["page"] == "" {
		params["page"] = "index"
	}
	layoutPath := path.Join(VIEWS_ROOT, "layout.html")
	contentPath := path.Join(VIEWS_ROOT, params["page"]+".html")

	if _, err := os.Stat(contentPath); err != nil {
		contentPath = path.Join(VIEWS_ROOT, "not-found.html")
	}

	contentTemplate, err := template.ParseFiles(layoutPath, contentPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := contentTemplate.Execute(w, params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc(STATIC_URL, processStatic)
	http.HandleFunc("/", processView)
	http.ListenAndServe(":8080", nil)
}
