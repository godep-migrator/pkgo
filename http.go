package pkgo

import (
	ag "github.com/subosito/anoa/github"
	"github.com/zenazn/goji/web"
	"net/http"
)

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)
	m.Get("/favicon.ico", FaviconHandler)
	m.Get("/robots.txt", RobotsHandler)
	m.Get("/css/theme.css", CSSHandler)
	m.Get("/css/auth-icons.png", AuthIconsHandler)
	m.Get("/about", AboutHandler)

	// oauth2
	m.Get("/auth", OauthHandler)
	m.Get("/auth/callback", OauthCallbackHandler)

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
	w.Write(box.MustBytes("stylesheets/auth-buttons.css"))
	w.Write(box.MustBytes("stylesheets/theme.css"))
}

func AuthIconsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("images/auth-icons.png"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tps["about"].Execute(w, nil)
}

func OauthHandler(w http.ResponseWriter, r *http.Request) {
	u := gh.AuthCodeURL("")
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func OauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")

	_, _, err := ag.Complete(gh, authCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
