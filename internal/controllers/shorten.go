package controllers

import (
	"fmt"
	"github.com/Francismensah/url-shortner-app/internal/url"
	"html/template"
	"net/http"
	"strings"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not implemented", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL not provide", http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		originalURL = "https://" + originalURL
	}

	hash, shortURL := url.Shorten(originalURL)

	fmt.Println(hash)

	data := map[string]string{
		"ShortURL": shortURL,
	}
	t, err := template.ParseFiles("internal/views/shorten.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
