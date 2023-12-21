package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ftqo/ftqo.dev/assets"
	"github.com/ftqo/ftqo.dev/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

var (
	tmpl *template.Template
	log  *slog.Logger
)

func main() {
	log = slog.New(slog.NewTextHandler(os.Stderr, nil))
	httplogger := httplog.NewLogger("http", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		MessageFieldName: "message",
		TimeFieldFormat:  time.RFC1123,
	})

	var err error
	tmpl, err = template.ParseFS(templates.T, "*/**.html")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(httplog.RequestLogger(httplogger))

	r.Mount("/assets", http.StripPrefix("/assets", http.FileServer(http.FS(assets.A))))

	counter := 0
	r.Post("/count", func(w http.ResponseWriter, r *http.Request) {
		counter++
		w.Write([]byte(strconv.Itoa(counter)))
	})
	r.Get("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strconv.Itoa(counter)))
	})

	r.Get("/*", serveTemplate)

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
	if !strings.HasSuffix(path, ".html") {
		path = path + ".html"
	}

	err := tmpl.ExecuteTemplate(w, "base.html", strings.TrimPrefix(path, "/"))
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
}
