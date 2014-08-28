package pkgo

import (
	"github.com/zenazn/goji/web"
	"io"
	"net/http"
)

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)

	return m
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "PKGO.ME")
}
