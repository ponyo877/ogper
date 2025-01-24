package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GenerateOGPPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stream, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer stream.Close()

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "title parameter is missing", http.StatusBadRequest)
		return
	}
	description := r.FormValue("description")
	if description == "" {
		http.Error(w, "description parameter is missing", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "name parameter is missing", http.StatusBadRequest)
		return
	}
	siteURL := r.FormValue("url")
	if siteURL == "" {
		http.Error(w, "url parameter is missing", http.StatusBadRequest)
		return
	}
	ogpPageURL, err := h.usecase.GenerateOGPPage(title, description, name, siteURL, stream, header.Size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", header.Filename)
	fmt.Fprintf(w, "URL: %s\n", ogpPageURL)
}

func (h *Handler) GetOGPPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	hash := r.PathValue("hash")
	html, err := h.usecase.GetOGPPage(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
