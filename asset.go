package pkgo

import (
	"net/http"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("favicon.ico"))
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("robots.txt"))
}

func AuthIconsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(box.MustBytes("images/auth-icons.png"))
}

func CSSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(box.MustBytes("stylesheets/pure-min.css"))
	w.Write(box.MustBytes("stylesheets/auth-buttons.css"))
	w.Write(box.MustBytes("stylesheets/theme.css"))
}
