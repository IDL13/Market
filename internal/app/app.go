package app

import (
	"fmt"
	http "net/http"
	"os"
	"time"

	handler "github.com/IDL13/Market/internal/handler/http"
)

type App struct {
	server    *http.Server
	handler   handler.Handler
	serverMux *http.ServeMux
}

func New() *App {
	app := &App{
		server: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	app.handler = handler.New()
	app.serverMux = http.NewServeMux()
	app.server.Handler = app.serverMux

	return app
}

func (a *App) Run(stop chan bool) {
	fmt.Println(`
╔══╗╔══╗─╔╗───╔╗╔══╗
╚╗╔╝║╔╗╚╗║║──╔╝║╚═╗║
─║║─║║╚╗║║║──╚╗║╔═╝║
─║║─║║─║║║║───║║╚═╗║
╔╝╚╗║╚═╝║║╚═╗─║║╔═╝║
╚══╝╚═══╝╚══╝─╚╝╚══╝`)

	fmt.Printf("[SERVER STARTED on %s]", a.server.Addr)
	fmt.Printf("[http://0.0.0.0:%s/docs/]", a.server.Addr)
	if err := a.server.ListenAndServe(); err != nil {
		stop <- true
		os.Exit(1)
	}
}
