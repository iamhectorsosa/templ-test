package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/a-h/templ"
	"github.com/iamhectorsosa/templ-test/templates"
)

//go:embed .dist/*
var dist embed.FS

func main() {
	mode := flag.String("mode", "", "Mode of operation: 'build' to build static files.")
	flag.Parse()

	switch *mode {
	case "build":
		build()
	case "preview":
		build()
		serve()
	default:
		serve()
	}
}

func serve() {
	homePage := templates.Home()
	subFS, err := fs.Sub(dist, ".dist")
	if err != nil {
		log.Fatalf("failed to strip prefix, err=%v", err)
	}

	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/css/", http.FileServer(http.FS(subFS)))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", pagesHandler)
}

func build() {
	rootPath := ".dist"
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	name := path.Join(rootPath, "index.html")
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file, err=%v", err)
	}

	if err = templates.Home().Render(context.Background(), file); err != nil {
		log.Fatalf("failed to write index page, err=%v", err)
	}

}

func preview() {
	subFS, err := fs.Sub(dist, ".dist")
	if err != nil {
		log.Fatalf("failed to strip prefix, err=%v", err)
	}

	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", http.FileServer(http.FS(subFS)))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", pagesHandler)
}
