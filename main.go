package main

import (
	"embed"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/wes-mil/templ-test/components"
)

//go:embed static/*
var static embed.FS

func main() {
	homePage := components.Index()

	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	if err := http.ListenAndServe(":3000", pagesHandler); err != nil {
		slog.Error("server failed", "error", err)
	}
}
