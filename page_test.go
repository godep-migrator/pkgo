package pkgo

import (
	"fmt"
	"github.com/zenazn/goji/web"
	check "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func init() {
	check.Suite(&PageSuite{suite: &BaseSuite{}})
}

type PageSuite struct {
	suite  *BaseSuite
	mux    *web.Mux
	server *httptest.Server
}

func (s *PageSuite) SetUpSuite(c *check.C) {
	s.suite.SetUpSuite(c)
	Initialize()

	s.mux = NewMux()
	s.server = httptest.NewServer(s.mux)
}

func (s *PageSuite) TearDownSuite(c *check.C) {
	s.server.Close()
	Terminate()
	s.suite.TearDownSuite(c)
}

func (s *PageSuite) TestHomeHandler(c *check.C) {
	res, err := http.Get(s.server.URL)
	c.Assert(err, check.IsNil)
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	c.Assert(err, check.IsNil)
	c.Assert(res.StatusCode, check.Equals, http.StatusOK)
	c.Assert(strings.Contains(string(b), "PKGO.ME"), check.Equals, true)
}

func (s *PageSuite) TestAboutHandler(c *check.C) {
	u := fmt.Sprintf("%s/about", s.server.URL)
	res, err := http.Get(u)
	c.Assert(err, check.IsNil)
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	c.Assert(err, check.IsNil)
	c.Assert(res.StatusCode, check.Equals, http.StatusOK)
	c.Assert(strings.Contains(string(b), "About"), check.Equals, true)
}
