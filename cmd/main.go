package main

import "github.com/IDL13/Market/internal/app"

func main() {
	stop := make(chan bool)

	entrypoint := app.New()

	entrypoint.Run(stop)

	<-stop
}
