package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ftqo/ftqo.dev/build"
	"github.com/ftqo/ftqo.dev/logger"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

type server struct {
	files map[string]*zip.File
	log   *slog.Logger
}

func main() {
	log := logger.GetLogger("http")
	s := server{log: log}

	fs, err := zip.NewReader(bytes.NewReader(build.F), int64(len(build.F)))
	if err != nil {
		log.Error("failed to read zipped file: " + err.Error())
		os.Exit(1)
	}

	// idea of snippet stolen from efron licht (see readme)
	s.files = make(map[string]*zip.File, len(fs.File))
	for _, f := range fs.File {
		s.files[strings.TrimPrefix(f.Name, "tmp/")] = f
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(slogchi.New(log))

	r.Get("/*", s.staticHandler)

	log.Info("listening on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err.Error())
	}
}

func (s server) staticHandler(w http.ResponseWriter, r *http.Request) {
	var f *zip.File
	path := strings.TrimPrefix(r.URL.Path, "/")

	switch {
	case path == "index.html":
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	case strings.HasSuffix(path, ".html"):
		http.Redirect(w, r, path[:len(path)-5], http.StatusMovedPermanently)
	default:
		if path == "" {
			path = "index.html"
		}
		if !strings.Contains(path, ".") {
			path = path + ".html"
		}

		var ok bool
		f, ok = s.files[path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	var body io.Reader
	var err error
	if strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") && f.Method == zip.Deflate {
		w.Header().Set("Content-Encoding", "deflate")
		body, err = f.OpenRaw()
	} else {
		body, err = f.Open()
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
