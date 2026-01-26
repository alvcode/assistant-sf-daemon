package app

import (
	"assistant-sf-daemon/internal/controller"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type App struct {
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(ctx context.Context) (App, error) {
	router := httprouter.New()

	return App{
		router: router,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx2 := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP(ctx2)
	})

	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {
	controllerInit := controller.New(a.router)
	errRoute := controllerInit.SetRoutes()
	if errRoute != nil {
		log.Fatalln(errRoute)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 8084))
	if err != nil {
		log.Fatalln(err)
	}

	c := cors.New(cors.Options{
		AllowedMethods:     strings.Split("GET,POST,PUT,PATCH,DELETE,OPTIONS", ","),
		AllowedOrigins:     strings.Split("*", ","),
		AllowedHeaders:     strings.Split("*", ","),
		AllowCredentials:   true,
		OptionsPassthrough: true,
		ExposedHeaders:     strings.Split("*", ","),
		Debug:              false,
	})

	hdl := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      hdl,
		WriteTimeout: time.Duration(15) * time.Minute,
		ReadTimeout:  time.Duration(15) * time.Minute,
		IdleTimeout:  time.Duration(15) * time.Minute,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
