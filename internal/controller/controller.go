package controller

import (
	"assistant-sf-daemon/internal/handler"
	"assistant-sf-daemon/internal/ucase"
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
	configHandler := handler.NewConfigHandler(ucase.NewConfigUseCase())

	/**
	запрос статуса
		проверяем наличия конфига. если его нет - возвращаем ошибку с кодом 1
	*/

	controller.router.Handler(
		http.MethodGet,
		"/api/heartbeat",
		handler.BuildHandler(heartbeatHandler.Heartbeat),
	)

	controller.router.Handler(
		http.MethodGet,
		"/api/config/status",
		handler.BuildHandler(configHandler.GetInitialStatus),
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
