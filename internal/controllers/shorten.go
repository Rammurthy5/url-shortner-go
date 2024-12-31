package controllers

import (
	"github.com/Rammurthy5/url-shortner-go/internal/urls"
	"html/template"
	"net/http"
	"strings"
)

func ShortenHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "url required", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	shortUrl := urls.Shorten(url)

	data := map[string]string{
		"ShortURL": shortUrl,
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
