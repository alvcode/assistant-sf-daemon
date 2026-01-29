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
	jobHandler := handler.NewJobHandler()

	controller.router.Handler(
		http.MethodGet,
		"/api/heartbeat",
		handler.BuildHandler(heartbeatHandler.Heartbeat),
	)

	controller.router.Handler(
		http.MethodPost,
		"/api/start-job",
		handler.BuildHandler(jobHandler.Start),
	)
	controller.router.Handler(
		http.MethodGet,
		"/api/status-job",
		handler.BuildHandler(jobHandler.GetStatus),
	)

	return nil
}
