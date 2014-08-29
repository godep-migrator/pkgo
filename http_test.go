package pkgo

import (
	check "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

func init() {
	check.Suite(&HandlerSuite{})
}

type HandlerSuite struct{}

func (s *HandlerSuite) TestHomeHandler(c *check.C) {
	m := NewMux()
	v := httptest.NewServer(m)

	res, err := http.Get(v.URL)
	c.Assert(err, check.IsNil)
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	c.Assert(err, check.IsNil)
	c.Assert(res.StatusCode, check.Equals, http.StatusOK)
	c.Assert(strings.Contains(string(b), "PKGO.ME"), check.Equals, true)
}
