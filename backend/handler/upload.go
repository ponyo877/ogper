package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var ogpPageDomain = os.Getenv("OGP_PAGE_DOMAIN")

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
	src, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer src.Close()

	name := r.FormValue("name")
	title := r.FormValue("title")
	description := r.FormValue("description")
	userHash := r.FormValue("user_hash")

	ogpPageURL, err := h.usecase.GenerateOGPPage(title, description, name, siteURL, userHash, src, header.Size)
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

func (h *Handler) ListSitesByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	userHash := r.URL.Query().Get("user_hash")
	sites, err := h.usecase.ListSitesByUserID(userHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Response struct {
		OGPerURL    string `json:"ogper_url"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Name        string `json:"name"`
		SiteURL     string `json:"site_url"`
		ImageURL    string `json:"image_url"`
	}
	responseList := []Response{}
	for _, site := range sites {
		response := Response{
			OGPerURL:    ogpPageDomain + "/" + site.Hash(),
			Title:       site.Title(),
			Description: site.Description(),
			Name:        site.Name(),
			SiteURL:     site.SiteURL(),
			ImageURL:    site.ImageURL(),
		}
		responseList = append(responseList, response)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseList)
}
