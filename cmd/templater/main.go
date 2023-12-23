package main

import (
	"bytes"
	"html/template"
	"os"
	"path"
	"path/filepath"

	"github.com/ftqo/ftqo.dev/logger"
)

func main() {
	log := logger.GetLogger("templater")
	input := "./templates"
	output := "./tmp"

	var files []string

	err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error(err.Error())
			panic(nil)
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		panic(nil)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Error(err.Error())
		panic(nil)
	}

	built := make(map[string][]byte)
	pages, err := os.ReadDir(path.Join(input, "pages"))
	if err != nil {
		log.Error(err.Error())
		panic(nil)
	}

	for _, p := range pages {
		b, err := os.ReadFile(path.Join(input, "pages", p.Name()))
		if err != nil {
			log.Error(err.Error())
			panic(nil)
		}

		var bb bytes.Buffer

		err = tmpl.ExecuteTemplate(&bb, "base.html", struct {
			Path    string
			Content template.HTML
		}{p.Name(), template.HTML(b)})
		if err != nil {
			log.Error(err.Error())
			panic(nil)
		}

		built[p.Name()] = bb.Bytes()
	}

	if err := os.MkdirAll(output, 0o777); err != nil {
		log.Error(err.Error())
		panic(nil)
	}

	for name, b := range built {
		fp := filepath.Join(output, name)
		err = os.WriteFile(fp, b, 0o666)
		if err != nil {
			log.Error(err.Error())
			panic(nil)
		}
		log.Info("built template", "path", fp)
	}
}
