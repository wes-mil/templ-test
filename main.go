package main

import (
	"embed"
	"log/slog"
	"net/http"

	"github.com/wes-mil/templ-test/components"
)

//go:embed static/*
var static embed.FS

func main() {
	pagesHandler := http.NewServeMux()
	pagesHandler.HandleFunc("/", getIndex)
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	pagesHandler.HandleFunc("POST /incrementCount", IncreaseCount)
	pagesHandler.HandleFunc("POST /decrementCount", DecrementCount)

	if err := http.ListenAndServe(":3000", pagesHandler); err != nil {
		slog.Error("server failed", "error", err)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	component := components.Index(Count)
	component.Render(r.Context(), w)
}
