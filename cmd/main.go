package main

import (
	"assistant-sf-daemon/internal/app"
	"assistant-sf-daemon/internal/locale"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	locale.InitLocales(ctx)
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	if err = a.Run(ctx); err != nil {
		log.Fatalln(err)
	}
}
