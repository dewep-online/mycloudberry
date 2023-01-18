package main

import (
	"github.com/dewep-online/goppy"
	"github.com/dewep-online/goppy/plugins/database"
	"github.com/dewep-online/goppy/plugins/http"
	"github.com/dewep-online/mycloudberry/app/components"
	"github.com/dewep-online/mycloudberry/app/core"
)

func main() {

	app := goppy.New()
	app.Plugins(
		http.WithHTTP(),
		http.WithHTTPDebug(),
		database.WithMySQL(),
		http.WithWebsocketServer(),
	)
	app.Plugins(core.Plugins()...)
	app.Plugins(components.Plugins()...)
	app.Run()

}
