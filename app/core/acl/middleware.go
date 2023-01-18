package acl

import (
	"fmt"
	"net/http"

	pluginshttp "github.com/dewep-online/goppy/plugins/http"
)

func BasicAuthMiddleware(skip map[string]struct{}) pluginshttp.Middleware {
	return func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			if _, ok := skip[r.URL.Path]; ok {
				next(w, r)
				return
			}
			u, p, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			fmt.Println(u, p)
			next(w, r)
		}
	}
}
