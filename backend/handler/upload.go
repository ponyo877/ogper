package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) GenerateOGPPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	siteURL := r.FormValue("url")
	if siteURL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}
	stream, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer stream.Close()
	name := r.FormValue("name")
	title := r.FormValue("title")
	description := r.FormValue("description")

	ogpPageURL, err := h.usecase.GenerateOGPPage(title, description, name, siteURL, stream, header.Size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Response struct {
		URL       string `json:"url"`
		CreatedAt string `json:"created_at"`
	}

	response := Response{
		URL:       ogpPageURL,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
