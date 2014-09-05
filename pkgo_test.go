package pkgo

import (
	"github.com/subosito/gotenv"
	check "gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type BaseSuite struct{}

func (s *BaseSuite) SetUpSuite(c *check.C) {
	gotenv.MustLoad(".env.test")
}

func (s *BaseSuite) TearDownSuite(c *check.C) {
	os.Clearenv()
}
