package pkgo

import (
	rice "github.com/GeertJohan/go.rice"
	"html/template"
)

var box *rice.Box
var tpl *template.Template

func init() {
	box = rice.MustFindBox("assets")

	s := box.MustString("templates/layout.html.tmpl")
	tpl = template.Must(template.New("_").Parse(s))
	tpl.Parse(box.MustString("templates/home.html.tmpl"))
}
