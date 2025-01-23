package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GenerateAltURL(w http.ResponseWriter, r *http.Request) {
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

	if err := h.usecase.GenerateAltURL(title, description, name, siteURL, stream, header.Size); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", header.Filename)
	fmt.Fprintf(w, "URL: %s\n", siteURL)
}

func (h *Handler) GetAltURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// <meta property="og:title" content="サイトタイトル" />
	// <meta property="og:description" content="サイトの説明" />
	// <meta property="og:type" content="website" />
	// <meta property="og:url" content="サイトURL" />
	// <meta property="og:image" content="サムネイル画像のURL" />
	// <meta property="og:site_name" content="サイト名" />
	// <meta name="twitter:card" content="summary_large_image" />
	// <meta name="twitter:title" content="サイトタイトル" />
	// <meta name="twitter:description" content="サイトの説明" />
	// <meta name="twitter:image" content="サムネイル画像のURL" />
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("og:title", "サイトタイトル")
	w.Header().Set("og:description", "サイトの説明")
	w.Header().Set("og:type", "website")
	w.Header().Set("og:url", "サイトURL")
	w.Header().Set("og:image", "サムネイル画像のURL")
	w.Header().Set("og:site_name", "サイト名")
	w.Header().Set("twitter:card", "summary_large_image")
	w.Header().Set("twitter:title", "サイトタイトル")
	w.Header().Set("twitter:description", "サイトの説明")
	w.Header().Set("twitter:image", "サムネイル画像のURL")
}
