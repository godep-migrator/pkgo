package pkgo

import (
	"github.com/subosito/gotenv"
	check "gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type BaseSuite struct {
	loadenv bool
}

func (s *BaseSuite) SetUpSuite(c *check.C) {
	tf := ".env.test"

	if s.isFileExist(tf) {
		err := gotenv.Load(tf)
		if err == nil {
			s.loadenv = true
		}
	}
}

func (s *BaseSuite) TearDownSuite(c *check.C) {
	if s.loadenv {
		os.Clearenv()
	}
}

func (s *BaseSuite) isFileExist(fn string) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}

	return true
}
