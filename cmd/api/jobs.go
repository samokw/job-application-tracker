package main

import (
	"fmt"
	"net/http"

	_ "github.com/julienschmidt/httprouter"
	"jobtracker.samokw.net/internal/data"
	"jobtracker.samokw.net/internal/validator"
)

func (app *application) createNewJobHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CompanyName     string      `json:"company_name"`
		Position        string      `json:"position"`
		Description     string      `json:"description,omitempty"`
		ApplicationDate string      `json:"appliaction_date"`
		Location        string      `json:"location"`
		Link            string      `json:"link,omitempty"`
		Status          data.Status `json:"status"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	job := &data.Job{
		CompanyName:     input.CompanyName,
		Position:        input.Position,
		Description:     input.Description,
		ApplicationDate: input.ApplicationDate,
		Location:        input.Location,
		Link:            input.Link,
		Status:          input.Status,
	}
	v := validator.New()
	if data.ValidateJob(v, job); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

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
		ApplicationDate: "2024-12-12",
		Location:        "Mountain View, CA",
		Link:            "https://jobs.google.com/software-engineer",
		Status:          data.Applied,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"job": job}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
