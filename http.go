package pkgo

import (
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
)

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)
	m.Get("/favicon.ico", FaviconHandler)

	return m
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	s := box.MustString("templates/home.html.tmpl")
	t := template.Must(template.New("home.html.tmpl").Parse(s))
	t.Execute(w, nil)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("favicon.ico"))
}
