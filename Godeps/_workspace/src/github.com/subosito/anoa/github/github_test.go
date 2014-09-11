package github

import (
	"github.com/golang/oauth2"
	"github.com/subosito/anoa"
	check "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

func init() {
	check.Suite(&GithubSuite{})
}

type mockTransport struct {
	rt func(req *http.Request) (resp *http.Response, err error)
}

func (t *mockTransport) RoundTrip(r *http.Request) (resp *http.Response, err error) {
	return t.rt(r)
}

type GithubSuite struct {
	options          *oauth2.Options
	defaultTransport http.RoundTripper
}

func (s *GithubSuite) SetUpSuite(c *check.C) {
	s.options = &oauth2.Options{
		ClientID:     "12345",
		ClientSecret: "abcdef",
		RedirectURL:  "http://example.com/callback",
		Scopes:       []string{"user:email"},
	}
}

func (s *GithubSuite) SetUpTest(c *check.C) {
	s.defaultTransport = http.DefaultTransport

	switch c.TestName() {
	case "GithubSuite.TestComplete":
		http.DefaultTransport = &mockTransport{
			rt: func(r *http.Request) (resp *http.Response, err error) {
				var body string

				switch r.URL.String() {
				case "https://github.com/login/oauth/access_token":
					body = fixtures["access_token"]
				case "https://api.github.com/user":
					body = fixtures["user"]
				}

				resp = &http.Response{}
				resp.StatusCode = http.StatusOK
				resp.Body = ioutil.NopCloser(strings.NewReader(body))
				return
			},
		}
	}
}

func (s *GithubSuite) TearDownTest(c *check.C) {
	http.DefaultTransport = s.defaultTransport
}

func (s *GithubSuite) TestNewConfig(c *check.C) {
	opts := s.options
	opts.Scopes = []string{"user:email", "public_repo", "gist"}

	cfg, err := NewConfig(opts)
	c.Assert(err, check.IsNil)

	u, err := url.Parse(cfg.AuthCodeURL(""))
	c.Assert(err, check.IsNil)

	c.Assert(u.Query().Get("client_id"), check.Equals, opts.ClientID)
	c.Assert(u.Query().Get("redirect_uri"), check.Equals, opts.RedirectURL)
	c.Assert(u.Query().Get("scope"), check.Equals, strings.Join(opts.Scopes, " "))
	c.Assert(u.Query().Get("response_type"), check.Equals, "code")

	u.RawQuery = ""
	c.Assert(u.String(), check.Equals, authURL)
}

func (s *GithubSuite) TestComplete(c *check.C) {
	cfg, err := NewConfig(s.options)
	c.Assert(err, check.IsNil)

	user, token, err := Complete(cfg, "1237890")
	c.Assert(err, check.IsNil)

	w := &anoa.User{
		Provider: "github",
		UID:      "1",
		Email:    "octocat@github.com",
		Nickname: "octocat",
		Name:     "monalisa octocat",
		Location: "San Francisco",
		Image:    "https://github.com/images/error/octocat_happy.gif",
		URLS: map[string]string{
			"github": "https://github.com/octocat",
			"blog":   "https://github.com/blog",
		},
		Extra: map[string]string(nil),
		Raw:   fixtures["user"],
	}

	c.Assert(user, check.DeepEquals, w)
	c.Assert(token, check.NotNil)
}

func (s *GithubSuite) TestNewClient(c *check.C) {
	cfg, err := NewConfig(s.options)
	c.Assert(err, check.IsNil)

	token := &oauth2.Token{AccessToken: "abc123"}
	client := NewClient(cfg, token)

	c.Assert(client, check.Not(check.DeepEquals), http.DefaultClient)
	c.Assert(client.Transport.(*oauth2.Transport).Token(), check.DeepEquals, token)
}

var fixtures = map[string]string{
	"user": `
		{
			"login": "octocat",
			"id": 1,
			"avatar_url": "https://github.com/images/error/octocat_happy.gif",
			"gravatar_id": "somehexcode",
			"url": "https://api.github.com/users/octocat",
			"html_url": "https://github.com/octocat",
			"followers_url": "https://api.github.com/users/octocat/followers",
			"following_url": "https://api.github.com/users/octocat/following{/other_user}",
			"gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
			"organizations_url": "https://api.github.com/users/octocat/orgs",
			"repos_url": "https://api.github.com/users/octocat/repos",
			"events_url": "https://api.github.com/users/octocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/octocat/received_events",
			"type": "User",
			"site_admin": false,
			"name": "monalisa octocat",
			"company": "GitHub",
			"blog": "https://github.com/blog",
			"location": "San Francisco",
			"email": "octocat@github.com",
			"hireable": false,
			"bio": "There once was...",
			"public_repos": 2,
			"public_gists": 1,
			"followers": 20,
			"following": 0,
			"created_at": "2008-01-14T04:33:35Z",
			"updated_at": "2008-01-14T04:33:35Z",
			"total_private_repos": 100,
			"owned_private_repos": 100,
			"private_gists": 81,
			"disk_usage": 10000,
			"collaborators": 8,
			"plan": {
				"name": "Medium",
				"space": 400,
				"private_repos": 20,
				"collaborators": 0
			}
		}
	`,
	"access_token": `
		{
			"access_token": "e72e16c7e42f292c6912e7710c838347ae178b4a",
			"scope": "user",
			"token_type": "bearer"
		}
	`,
}
