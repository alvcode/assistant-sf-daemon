package controller

import (
	"assistant-sf-daemon/internal/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Init struct {
	router *httprouter.Router
}

func New(router *httprouter.Router) *Init {
	return &Init{
		router: router,
	}
}

func (controller *Init) SetRoutes() error {
	//repos := repository.NewRepositories(controller.cfg, controller.db, controller.minio)

	//handler.InitHandler(repos, controller.cfg)

	controller.router.NotFound = handler.BuildHandler(handler.PageNotFoundHandler)
	controller.router.MethodNotAllowed = handler.BuildHandler(handler.PageNotFoundHandler)

	heartbeatHandler := handler.NewHeartbeatHandler()

	controller.router.Handler(
		http.MethodGet,
		"/api/heartbeat",
		handler.BuildHandler(heartbeatHandler.Heartbeat),
	)

	return nil
}
