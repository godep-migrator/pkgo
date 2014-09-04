package pkgo

import (
	"encoding/gob"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/golang/oauth2"
	"github.com/gorilla/sessions"
	"github.com/subosito/anoa"
	ag "github.com/subosito/anoa/github"
	redistore "gopkg.in/boj/redistore.v1"
	"html/template"
	"net/http"
	"os"
)

var box *rice.Box
var gh *oauth2.Config
var sst *redistore.RediStore
var tps = map[string]*template.Template{}

func Initialize() {
	initAssets()
	initOauth2()
	initSessions()
}

func initAssets() {
	box = rice.MustFindBox("assets")
	base := box.MustString("templates/layout.html.tmpl")
	tmps := []string{"home", "about"}

	for i := range tmps {
		f := tmps[i]
		t := template.Must(template.New("_").Parse(base))

		p, err := t.Parse(box.MustString(fmt.Sprintf("templates/%s.html.tmpl", f)))
		if err != nil {
			panic(fmt.Errorf("Error parsing template for %s: %s", f, err))
		}

		tps[f] = p
	}
}

func initOauth2() {
	var err error

	authConfig := &oauth2.Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("CLIENT_REDIRECT_URL"),
		Scopes:       []string{"user:email"},
	}

	gh, err = ag.NewConfig(authConfig)
	if err != nil {
		panic(err)
	}
}

func initSessions() {
	var err error

	sst, err = redistore.NewRediStore(10, "tcp", os.Getenv("REDIS_URL"), "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}

	// register session objects
	gob.Register(&oauth2.Token{})
	gob.Register(&anoa.User{})
}

func Terminate() {
	sst.Close()
}

func session(r *http.Request) *sessions.Session {
	ses, err := sst.Get(r, "pkgo-session")
	if err != nil {
		panic(err)
	}

	return ses
}
