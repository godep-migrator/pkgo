package pkgo

import (
	"github.com/subosito/anoa"
	"github.com/zenazn/goji/web"
	"net/http"
)

func HomeHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	tps["home"].Execute(w, TemplateData{User: c.Env["current-user"].(*anoa.User)})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tps["about"].Execute(w, nil)
}
