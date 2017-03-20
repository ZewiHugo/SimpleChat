package routers

import (
	"net/http"
	"SimpleChat/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var privateRoutes = Routes{
	Route{"GET",	"/users",			handlers.UserList,},
	Route{"GET",	"/users/{id}",		handlers.UserGet,},
	Route{"POST",	"/users",			handlers.UserCreate,},
	Route{"DELETE",	"/users/{id}",		handlers.UserDelete,},
	Route{"DELETE",	"/users",			handlers.UserDeleteAll,},
	Route{"POST",	"/users/verify",	handlers.VerifyUser,},
}
