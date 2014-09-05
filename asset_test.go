package pkgo

import (
	"bytes"
	"github.com/zenazn/goji/web"
	check "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func init() {
	check.Suite(&AssetSuite{suite: &BaseSuite{}})
}

type AssetSuite struct {
	suite  *BaseSuite
	mux    *web.Mux
	server *httptest.Server
}

func (s *AssetSuite) SetUpSuite(c *check.C) {
	s.suite.SetUpSuite(c)
	Initialize()
}

func (s *AssetSuite) SetUpTest(c *check.C) {
	s.mux = web.New()
	s.server = httptest.NewServer(s.mux)
}

func (s *AssetSuite) TearDownTest(c *check.C) {
	s.server.Close()
}

func (s *AssetSuite) TearDownSuite(c *check.C) {
	Terminate()
	s.suite.TearDownSuite(c)
}

func (s *AssetSuite) TestFaviconHandler(c *check.C) {
	s.mux.Get("/", FaviconHandler)
	s.assertR(c, s.server.URL, "image/vnd.microsoft.icon", []string{"assets/favicon.ico"})
}

func (s *AssetSuite) TestRobotsHandler(c *check.C) {
	s.mux.Get("/", RobotsHandler)
	s.assertR(c, s.server.URL, "text/plain; charset=utf-8", []string{"assets/robots.txt"})
}

func (s *AssetSuite) TestAuthIconsHandler(c *check.C) {
	s.mux.Get("/", AuthIconsHandler)
	s.assertR(c, s.server.URL, "image/png", []string{"assets/images/auth-icons.png"})
}

func (s *AssetSuite) TestCSSHandler(c *check.C) {
	s.mux.Get("/", CSSHandler)
	s.assertR(c, s.server.URL, "text/css", []string{
		"assets/stylesheets/pure-min.css",
		"assets/stylesheets/auth-buttons.css",
		"assets/stylesheets/theme.css",
	})
}

func (s *AssetSuite) assertR(c *check.C, u, ctype string, fnames []string) {
	res, err := http.Get(u)
	c.Assert(err, check.IsNil)
	defer res.Body.Close()

	c.Assert(res.StatusCode, check.Equals, http.StatusOK)
	c.Assert(res.Header.Get("Content-Type"), check.Equals, ctype)

	b, err := ioutil.ReadAll(res.Body)
	c.Assert(err, check.IsNil)

	w := bytes.NewBuffer([]byte{})

	for _, fn := range fnames {
		d, err := ioutil.ReadFile(fn)
		c.Assert(err, check.IsNil)

		_, err = w.Write(d)
		c.Assert(err, check.IsNil)
	}

	c.Assert(len(b), check.Equals, w.Len())
	c.Assert(b, check.DeepEquals, w.Bytes())
}
