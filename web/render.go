package web

import (
	"html/template"
	"net/http"

	tmpl "github.com/hi20160616/udp2mysql/web/templates"
)

type Page struct {
	Title string
	Data  interface{}
}

var tmpls = template.New("")

func init() {
	tmpls.Funcs(template.FuncMap{
		"increaseOne": increaseOne,
	})
	tmpls = template.Must(tmpls.ParseFS(tmpl.FS, "default/*.html"))
}

func derive(w http.ResponseWriter, tmpl string, p *Page) error {
	return tmpls.ExecuteTemplate(w, tmpl+".html", p)
}

func increaseOne(x int) int {
	return x + 1
}
