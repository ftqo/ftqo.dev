package generate

//go:generate tailwindcss -i ./templates/tailwind.css -o ./tmp/assets/css/styles.css --minify
//go:generate go run cmd/templater/main.go
//go:generate go run cmd/zipper/main.go
