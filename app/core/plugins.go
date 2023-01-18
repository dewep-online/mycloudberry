package core

import (
	"github.com/dewep-online/goppy/plugins"
)

func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		New(),
	}
}
