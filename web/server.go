package web

import (
	"net/http"

	tmpl "github.com/hi20160616/udp2mysql/web/templates"
)

type WebServer struct {
	http.Server
}

func NewWebServer(addr string) *WebServer {
	return &WebServer{http.Server{
		Addr:    addr,
		Handler: getHandler(),
	}}
}

func getHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		homeHandler(w, r)
	})
	mux.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.FS(tmpl.FS))))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	derive(w, "home", &Page{"Home", "just home content"})
}
