package controllers

import (
	"fmt"
	"github.com/Rammurthy5/url-shortner-go/internal/utils"
	"github.com/Rammurthy5/url-shortner-go/internal/validators"
	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
	"html/template"
	"net/http"
	"strings"
)

// ShortenRequest is used to capture form data.
type ShortenRequest struct {
	URL string `form:"url" validate:"required,url"`
}

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

	// Parse the form data
	decoder := form.NewDecoder()
	var req ShortenRequest

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Decode form data into the struct
	if err := decoder.Decode(&req, r.Form); err != nil {
		http.Error(w, "Failed to decode form", http.StatusBadRequest)
		return
	}

	// Validate the URL
	if err := validators.ValidateURL(req.URL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the struct using go-playground/validator
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	// If the URL doesn't start with http:// or https://, add https:// as default
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		req.URL = "https://" + req.URL
	}

	shortURL := utils.FetchShortURL(c.Db, req.URL)
	if shortURL != "" {
		_log.Info("Url fetched from DB")
	}
	if shortURL == "" {
		shortURL = utils.Shorten(req.URL)
		err := utils.StoreShortURL(c.Db, req.URL, shortURL)
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
