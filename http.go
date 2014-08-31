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

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tps["about"].Execute(w, nil)
}
