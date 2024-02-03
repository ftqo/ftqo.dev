package generate

//go:generate go run cmd/templater/main.go
//go:generate npm run tailwind
// tailwind must be generated before zipper... zips the tailwind output

//go:generate go run cmd/zipper/main.go
