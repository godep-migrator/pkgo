package pkgo

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"html/template"
)

var box *rice.Box
var tps = map[string]*template.Template{}

func init() {
	box = rice.MustFindBox("assets")

	s := box.MustString("templates/layout.html.tmpl")

	templates := []string{
		"home",
		"about",
	}

	for _, f := range templates {
		t := template.Must(template.New("_").Parse(s))

		p, err := t.Parse(box.MustString(fmt.Sprintf("templates/%s.html.tmpl", f)))
		if err != nil {
			panic(fmt.Errorf("Error parsing template for %s: %s", f, err))
		}

		tps[f] = p
	}
}
