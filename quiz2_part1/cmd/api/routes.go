//Filename:cmd/api/routes.go
package main

import	(
	"net/http"

	"github.com/julienschmidt/httprouter"
)
func (app *application) routes() *httprouter.Router{
	//Create new httprouter instance
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.createEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.showEntryHandler)
	return router
}

