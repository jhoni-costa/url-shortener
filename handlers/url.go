package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jhoni-costa/url-shortener/storage"
)

// URLHandler depends on our storage to function ( Dependency Injection)
type URLHandler struct {
	Storage *storage.MapStorage
}

// RequestBody defines the JSON structure of the request
type RequestBody struct {
	URL string `json:"url"`
}

// Shorten handles the POST to create the short URL
func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	// Decode the JSON from the request body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		http.Error(w, "JSON inválido ou campo 'url' vazio", http.StatusBadRequest)
		return
	}

	// Generate a random code of 6 characters
	code := generateCode(6)

	// Save in memory storage
	h.Storage.Save(code, body.URL)

	// Respond with the generated code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", code),
	})
}

// Redirect handles the GET and redirects the user
func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Get the code from the URL (ex: /abcde -> abcde)
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.Error(w, "Código inválido", http.StatusBadRequest)
		return
	}

	originalURL, err := h.Storage.Get(code)
	if err != nil {
		http.Error(w, "URL não encontrada", http.StatusNotFound)
		return
	}

	// Redirects the user (Status 302 Found)
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// Utility function to generate random string
func generateCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
