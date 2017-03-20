package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"SimpleChat/middlewares"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	commonMiddleware := negroni.New(
		negroni.HandlerFunc(middlewares.Logger),
	)

	for _, route := range privateRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(commonMiddleware.With(
				negroni.HandlerFunc(middlewares.RequireAuthentication),
				negroni.Wrap(route.HandlerFunc),
			))
	}

	for _, route := range publicRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(commonMiddleware.With(negroni.Wrap(route.HandlerFunc)))
	}

	return router
}