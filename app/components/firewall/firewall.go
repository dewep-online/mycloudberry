package firewall

import (
	"net"

	"github.com/dewep-online/goppy/plugins"
	"github.com/dewep-online/goppy/plugins/http"
)

func New() plugins.Plugin {
	return plugins.Plugin{
		Inject: func(conf *http.Config) *UFW {
			rules := make([]Rule, 0)
			for _, config := range conf.Config {
				if alias, ok := networkType[config.Network]; !ok {
					_, p, err := net.SplitHostPort(config.Addr)
					if err != nil {
						continue
					}
					rules = append(rules, Rule{
						Port:    str2int(p),
						Network: alias,
					})
				}
			}

			return newUfw(rules)
		},
		Resolve: func(routes http.RouterPool, obj *UFW) {
			router := routes.Main()
			router.Get("/api/firewall/status", func(ctx http.Ctx) {
				ctx.SetBody(205)

				obj.Status()
			})
		},
	}
}
