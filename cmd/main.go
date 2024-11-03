package main

import (
	"context"
	"os"

	"github.com/IDL13/Market/internal/app"
	"github.com/IDL13/Market/pkg/prometheus"
	"go.uber.org/fx"
)

func main() {
	fx.New(CreateApp()).Run()
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			app.NewApp,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, app *app.App) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						stop := make(chan bool)
						go func() {
							err := prometheus.Listen("0.0.0.0:8082")
							if err != nil {
								stop <- true
							}
						}()
						app.Run(stop)
						<-stop
						os.Exit(1)
						return nil
					},
					OnStop: func(context.Context) error {
						return nil
					},
				})
			},
		),
	)
}
