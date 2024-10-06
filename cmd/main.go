package main

import (
	"os"

	"github.com/IDL13/Market/internal/app"
	"github.com/IDL13/Market/pkg/prometheus"
)

func main() {
	stop := make(chan bool)

	entrypoint := app.New()

	go func() {
		err := prometheus.Listen("0.0.0.0:8082")
		if err != nil {
			stop <- true
		}
	}()

	entrypoint.Run(stop)

	<-stop
	os.Exit(1)
}
