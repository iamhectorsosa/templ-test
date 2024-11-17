package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/iamhectorsosa/templ-test/templates"
)

func main() {
	mode := flag.String("mode", "", "Mode of operation: 'build' to build static files.")
	flag.Parse()

	switch *mode {
	case "build":
		build()
	default:
		build()
		serve()
	}
}

func build() {
	rootPath := ".dist"
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		log.Fatalf("failed to create output directory, err=%v", err)
	}

	name := path.Join(rootPath, "index.html")
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to create output file, err=%v", err)
	}

	if err = templates.Home().Render(context.Background(), file); err != nil {
		log.Fatalf("failed to write index page, err=%v", err)
	}
}

func serve() {
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", http.FileServer(http.Dir(".dist")))

	fmt.Println("Listening on http://localhost:3000/")
	http.ListenAndServe(":3000", pagesHandler)
}
