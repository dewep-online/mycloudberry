package core

import (
	"github.com/dewep-online/goppy/plugins"
	"github.com/dewep-online/goppy/plugins/http"
	"github.com/dewep-online/mycloudberry/app/core/acl"
	"github.com/dewep-online/mycloudberry/app/core/web"
)

func New() plugins.Plugin {
	return plugins.Plugin{
		Resolve: func(routes http.RouterPool, ws http.WebsocketServer) {
			router := routes.Main()
			router.Use(http.ThrottlingMiddleware(100))
			router.Use(web.StaticMiddleware())
			router.Use(acl.BasicAuthMiddleware(map[string]struct{}{
				"/health": {},
			}))

			router.Get("/ws", ws.Handling)
			router.Get("/health", func(ctx http.Ctx) {
				ctx.SetBody(200).String("ok")
			})
		},
	}
}
