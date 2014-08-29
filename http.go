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

	return m
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("favicon.ico"))
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("robots.txt"))
}
