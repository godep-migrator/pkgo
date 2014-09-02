package pkgo

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)
	m.Get("/favicon.ico", FaviconHandler)
	m.Get("/robots.txt", RobotsHandler)
	m.Get("/css/theme.css", CSSHandler)
	m.Get("/about", AboutHandler)

	return m
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tps["home"].Execute(w, nil)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("favicon.ico"))
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("robots.txt"))
}

func CSSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(box.MustBytes("stylesheets/pure-min.css"))
	w.Write(box.MustBytes("stylesheets/theme.css"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tps["about"].Execute(w, nil)
}
