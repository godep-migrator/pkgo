package pkgo

import (
	rice "github.com/GeertJohan/go.rice"
)

var box *rice.Box

func Boot() {
	box = rice.MustFindBox("assets")
}
