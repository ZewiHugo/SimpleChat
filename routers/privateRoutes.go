package routers

import (
	"SimpleChat/handlers"
)

var publicRoutes = Routes{
	Route{"GET",	"/",				handlers.Index,},
	Route{"POST",	"/login",			handlers.Login,},
}
