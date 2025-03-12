package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ppastene/go-shortener/internal/config"
	"github.com/ppastene/go-shortener/internal/services"
	"github.com/ppastene/go-shortener/pkg/keygen"
)

type ShortenerController struct {
	ShortenerService *services.ShortenerService
	cfg              *config.Config
	keygen           *keygen.Keygen
}

func NewShortenerController(service services.ShortenerService, keygen keygen.Keygen, cfg config.Config) *ShortenerController {
	return &ShortenerController{&service, &cfg, &keygen}
}

func (sc ShortenerController) RedirectUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
		return
	}
	key := r.URL.Path[len("/redirect/"):]
	if key == "" {
		http.Error(w, "Shortkey is missing", http.StatusBadRequest)
		return
	}
	shortenedUrl, err := sc.ShortenerService.GetUrl(key)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, shortenedUrl.Url, http.StatusFound)
}

func (sc ShortenerController) SaveUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
		return
	}
	// Validate if
	// the url field is not empty
	// the url is a valid one
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "Missing URL from input", http.StatusBadRequest)
		return
	}
	// Make this a loop, so if the shortcode exists generate other and try again
	shortcode, err := sc.keygen.Generate(uint(sc.cfg.KeyLength))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shortenedUrl := fmt.Sprintf("http://localhost:8080/redirect/%s", shortcode)
	err = sc.ShortenerService.SaveUrl(shortcode, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl, err := template.ParseFiles("views/layout/_header.html", "views/short.html", "views/layout/_footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "short", shortenedUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (sc ShortenerController) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/layout/_header.html", "views/index.html", "views/layout/_footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (sc ShortenerController) Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("views/layout/_header.html", "views/error.html", "views/layout/_footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "error", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (sc ShortenerController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("views/layout/_header.html", "views/list.html", "views/layout/_footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "list", sc.ShortenerService.ListShortcodes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
