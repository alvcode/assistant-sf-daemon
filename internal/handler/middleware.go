package handler

import (
	"assistant-sf-daemon/internal/locale"
	"net/http"
)

const (
	LocaleMW = "LocaleMW"
)

var MapMiddleware = map[string]Middleware{
	LocaleMW: LocaleMiddleware,
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func BuildHandler(h http.HandlerFunc, middlewares ...string) http.HandlerFunc {
	stack := []Middleware{
		MapMiddleware[LocaleMW],
	}

	for _, mw := range middlewares {
		if middlewareFunc, exists := MapMiddleware[mw]; exists {
			stack = append(stack, middlewareFunc)
		}
	}
	return ApplyMiddleware(h, stack...)
}

func LocaleMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return locale.Middleware(next)
}
