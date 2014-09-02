package pkgo

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/golang/oauth2"
	"html/template"
	"os"
)

var box *rice.Box
var gh *oauth2.Config
var tps = map[string]*template.Template{}

func Initialize() {
	initAssets()
	initOauth2()
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

	authURL := "https://github.com/login/oauth/authorize"
	authToken := "https://github.com/login/oauth/access_token"
	authConfig := &oauth2.Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("CLIENT_REDIRECT_URL"),
		Scopes:       []string{"email"},
	}

	gh, err = oauth2.NewConfig(authConfig, authURL, authToken)
	if err != nil {
		panic(err)
	}
}
