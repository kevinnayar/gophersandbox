package router

import (
	"github.com/kevinnayar/gophersandbox/cmd/api/handlers/roles"
	"github.com/kevinnayar/gophersandbox/cmd/api/handlers/users"
	"github.com/kevinnayar/gophersandbox/pkg/application"

	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/roles", roles.GetAll(app))
	mux.GET("/users", users.GetAll(app))

	return mux
}
