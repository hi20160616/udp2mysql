package web

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(list|search|udp|del|)/(.*?)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, *Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
		}
		fn(w, r, &Page{})
	}
}
