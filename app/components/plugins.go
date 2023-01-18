package components

import (
	"github.com/dewep-online/goppy/plugins"
	"github.com/dewep-online/mycloudberry/app/components/firewall"
)

func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		firewall.New(),
	}
}
