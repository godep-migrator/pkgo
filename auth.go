package pkgo

import (
	ag "github.com/subosito/anoa/github"
	"net/http"
)

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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cs := session(r)
	cs.Options.MaxAge = -1
	cs.Save(r, w)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
