package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ftqo/ftqo.dev/assets"
	"github.com/ftqo/ftqo.dev/templates"
	"github.com/go-chi/chi/v5"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFS(templates.T, "*/**.html")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	r := chi.NewRouter()

	r.Mount("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.A))))

	r.Get("/", serveTemplate)

	counter := 0
	r.Post("/counter", func(w http.ResponseWriter, r *http.Request) {
		counter++
		w.Write([]byte(strconv.Itoa(counter)))
	})
	r.Get("/counter", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strconv.Itoa(counter)))
	})

	err = http.ListenAndServe(":8080", r)
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
