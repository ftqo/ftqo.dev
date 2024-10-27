package main

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ftqo/ftqo.dev/logger"
)

func main() {
	input := "tmp"
	output := "build/files.zip"
	assets := "assets"
	log := logger.GetLogger("zipper")

	as, err := filepath.Glob(filepath.Join(assets, "*"))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	args := append([]string{"-R"}, as...)
	args = append(args, filepath.Join(input, "assets"))

	cmd := exec.Command("cp", args...)
	err = cmd.Run()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	zipFile, err := os.Create(output)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(input, func(filePath string, d os.FileInfo, err error) error {
		if strings.Contains(d.Name(), ".zip") || strings.Contains(d.Name(), ".gz") || strings.Contains(d.Name(), ".go") || d.IsDir() || strings.Contains(filePath, "/bin") {
			return nil
		}

		if filePath == output {
			return nil
		}
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(d)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(strings.Replace(filePath, input, "", 1), string(filepath.Separator))
		header.Name = strings.TrimLeft(header.Name, string(filepath.Separator))
		header.Name = strings.ReplaceAll(header.Name, string(filepath.Separator), "/")

		// HTML files should be in the root
		header.Name = strings.Replace(header.Name, "html/", "", 1)

		if d.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !d.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			return err
		}

		return nil
	})
	if err != nil {
		log.Error(err.Error())
		panic(nil) // run deferred functions
	}

	log.Info("built zip", "path", filepath.Clean(output))
}
