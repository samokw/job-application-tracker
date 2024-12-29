package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/julienschmidt/httprouter"
	"jobtracker.samokw.net/internal/data"
)

func (app *application) createNewJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "craete a new movie")
}

func (app *application) showJobHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	job := data.Job{
		ID:              id,
		CompanyName:     "Google",
		Position:        "Software Engineer",
		Description:     "Develop and maintain high-quality software.",
		ApplicationDate: time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC),
		Location:        "Mountain View, CA",
		Link:            "https://jobs.google.com/software-engineer",
		Status:          data.Applied,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"job":job}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
