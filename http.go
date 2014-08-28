package pkgo

import (
	"io"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "PKGO.ME")
}
