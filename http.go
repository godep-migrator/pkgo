package pkgo

import (
	"github.com/subosito/anoa"
	ag "github.com/subosito/anoa/github"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
)

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)
	m.Get("/about", AboutHandler)
	m.Get("/logout", LogoutHandler)

	// assets
	m.Get("/robots.txt", RobotsHandler)
	m.Get("/favicon.ico", FaviconHandler)
	m.Get("/css/theme.css", CSSHandler)
	m.Get("/css/auth-icons.png", AuthIconsHandler)

	// oauth2
	m.Get("/auth", OauthHandler)
	m.Get("/auth/callback", OauthCallbackHandler)

	// middlewares
	m.Use(middleware.EnvInit)
	m.Use(AuthMiddleware)

	return m
}

type TemplateData struct {
	User *anoa.User
}

func HomeHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	tps["home"].Execute(w, TemplateData{User: c.Env["current-user"].(*anoa.User)})
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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cs := session(r)
	cs.Options.MaxAge = -1
	cs.Save(r, w)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func OauthHandler(w http.ResponseWriter, r *http.Request) {
	u := gh.AuthCodeURL("")
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func OauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")

	user, token, err := ag.Complete(gh, authCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	cs := session(r)
	cs.Values["current-user"] = user
	cs.Values["access-token"] = token
	cs.Save(r, w)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
