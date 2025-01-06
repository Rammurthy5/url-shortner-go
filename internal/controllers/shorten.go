package controllers

import (
	"fmt"
	"github.com/Rammurthy5/url-shortner-go/internal/utils"
	"html/template"
	"net/http"
	"strings"
)

type ShortenController struct {
	*BaseController
}

func NewShortenController(base *BaseController) *ShortenController {
	return &ShortenController{BaseController: base}
}

func (c *ShortenController) ServeHandle(w http.ResponseWriter, r *http.Request) {
	_log := c.Log.Named("shortenController")
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
		url = "https://" + url
	}
	shortURL := utils.FetchShortURL(c.Db, url)
	if shortURL != "" {
		_log.Info("Url fetched from DB")
	}
	if shortURL == "" {
		shortURL = utils.Shorten(url)
		err := utils.StoreShortURL(c.Db, url, shortURL)
		if err != nil {
			_log.Error(fmt.Sprintf("URL store to database has failed %s", err))
		}
	}

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
