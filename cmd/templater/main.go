package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ftqo/ftqo.dev/logger"
	"golang.org/x/exp/maps"
)

func main() {
	log := logger.GetLogger("templater")
	input := "templates"
	output := "tmp"

	var files []string

	err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	built := make(map[string][]byte)
	pages, err := os.ReadDir(path.Join(input, "pages"))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	for _, p := range pages {
		b, err := os.ReadFile(path.Join(input, "pages", p.Name()))
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		var bb bytes.Buffer

		if _, ok := strings.CutPrefix(p.Name(), "_"); ok {
			built[strings.TrimPrefix(p.Name(), "_")] = b
			continue
		}

		sPath, _ := strings.CutSuffix(p.Name(), ".html")

		err = tmpl.ExecuteTemplate(&bb, "base.html", struct {
			Path    string
			Content template.HTML
		}{sPath, template.HTML(b)})
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		built[p.Name()] = bb.Bytes()
	}

	log.Debug(fmt.Sprintf("built %+v", maps.Keys(built)))

	err = os.RemoveAll(output)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	if err := os.MkdirAll(output, 0o755); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	for name, b := range built {
		fp := filepath.Join(output, name)
		err = os.WriteFile(fp, b, 0o644)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		log.Info("built template", "path", fp)
	}
}
