package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/ftqo/ftqo.dev/assets"
	"github.com/ftqo/ftqo.dev/templates"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFS(templates.T, "*/**.html")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.A))))
	http.HandleFunc("/", serveTemplate)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error(err.Error())
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	if strings.HasSuffix(path, ".html") {
		err := tmpl.ExecuteTemplate(w, strings.TrimPrefix(path, "/"), nil)
		if err != nil {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}
