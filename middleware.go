package pkgo

import (
	"github.com/golang/oauth2"
	"github.com/subosito/anoa"
	"github.com/zenazn/goji/web"
	"net/http"
)

func AuthMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cs := session(r)

		var user *anoa.User
		var token *oauth2.Token

		if u, ok := cs.Values["current-user"]; ok {
			user = u.(*anoa.User)
		}

		if t, ok := cs.Values["access-token"]; ok {
			token = t.(*oauth2.Token)
		}

		c.Env["current-user"] = user
		c.Env["access-token"] = token

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
