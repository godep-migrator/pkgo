package pkgo

import (
	"github.com/google/go-github/github"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)
	m.Get("/favicon.ico", FaviconHandler)
	m.Get("/robots.txt", RobotsHandler)
	m.Get("/css/theme.css", CSSHandler)
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
	w.Write(box.MustBytes("stylesheets/theme.css"))
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
	log.Println(authCode)

	token, err := gh.Exchange(authCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t := gh.NewTransport()
	t.SetToken(token)

	c := github.NewClient(&http.Client{Transport: t})
	u, _, err := c.Users.Get("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(u.Email)
}
