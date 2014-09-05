package pkgo

import (
	"github.com/subosito/gotenv"
	check "gopkg.in/check.v1"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type BaseSuite struct{}

func (s *BaseSuite) SetUpSuite(c *check.C) {
	err := gotenv.Load(".env.test")
	if err != nil {
		log.Println(err)
	}
}

func (s *BaseSuite) TearDownSuite(c *check.C) {
	os.Clearenv()
}
