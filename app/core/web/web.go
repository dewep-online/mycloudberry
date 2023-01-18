package web

import (
	"net/http"

	pluginshttp "github.com/dewep-online/goppy/plugins/http"
	"github.com/deweppro/go-static"
)

//go:generate static ./../../../web/dist/application ui

var ui static.Reader

func StaticMiddleware() pluginshttp.Middleware {
	list := ui.List()
	files := make(map[string]struct{})
	for _, uri := range list {
		files[uri] = struct{}{}
	}

	return func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			uri := r.URL.Path
			switch uri {
			case "", "/":
				uri = "/index.html"
			}
			if _, ok := files[uri]; ok {
				if err := ui.ResponseWrite(w, uri); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				return
			}
			next(w, r)
		}
	}
}
