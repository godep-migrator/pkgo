package github

import (
	"fmt"
	"github.com/golang/oauth2"
	"github.com/stretchr/objx"
	"github.com/subosito/anoa"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	provider string = "github"
	endpoint string = "https://api.github.com"
	authURL  string = "https://github.com/login/oauth/authorize"
	tokenURL string = "https://github.com/login/oauth/access_token"
)

func NewConfig(opts *oauth2.Options) (*oauth2.Config, error) {
	return oauth2.NewConfig(opts, authURL, tokenURL)
}

func Complete(cfg *oauth2.Config, code string) (user *anoa.User, token *oauth2.Token, err error) {
	token, err = cfg.Exchange(code)
	if err != nil {
		return
	}

	u := relPath("user")
	c := NewClient(cfg, token)
	resp, err := c.Get(u)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	user, err = constructUser(string(b))
	return
}

func NewClient(cfg *oauth2.Config, token *oauth2.Token) *http.Client {
	t := cfg.NewTransport()
	t.SetToken(token)

	return &http.Client{Transport: t}
}

func relPath(s string) string {
	return fmt.Sprintf("%s/%s", endpoint, s)
}

func constructUser(str string) (*anoa.User, error) {
	m, err := objx.FromJSON(str)
	if err != nil {
		return nil, err
	}

	user := &anoa.User{}
	user.Provider = provider
	user.UID = strconv.FormatFloat(m.Get("id").Float64(), 'f', 0, 64)
	user.Email = m.Get("email").Str()
	user.Nickname = m.Get("login").Str()
	user.Name = m.Get("name").Str()
	user.Location = m.Get("location").Str()
	user.Image = m.Get("avatar_url").Str()
	user.Raw = str
	user.URLS = map[string]string{
		"github": m.Get("html_url").Str(),
		"blog":   m.Get("blog").Str(),
	}

	return user, nil
}
