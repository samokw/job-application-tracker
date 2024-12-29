package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/jobs", app.createNewJobHandler)
	router.HandlerFunc(http.MethodGet, "/jobs/:id", app.showJobHandler)

	return router
}
