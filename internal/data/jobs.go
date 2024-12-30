package data

import "jobtracker.samokw.net/internal/validator"

type Status int

const (
	Saved Status = iota
	Applied
	Interview
	Accepted
	Rejected
)

type Job struct {
	ID              int64  `json:"id"`
	CompanyName     string `json:"company_name"`
	Position        string `json:"position"`
	Description     string `json:"description,omitempty"`
	ApplicationDate string `json:"appliaction_date"`
	Location        string `json:"location"`
	Link            string `json:"link,omitempty"`
	Status          Status `json:"status"`
}
func ValidateJob(v *validator.Validator, job *Job){
	v.Check(job.CompanyName != "", "company_name", "must be provided")
	v.Check(len(job.CompanyName) <= 500, "company_name", "must not be more than 500 bytes long")
	v.Check(job.Position != "", "position", "must be provided")
	v.Check(len(job.Position) <= 500, "position", "must not be more than 500 bytes long")
	v.Check(job.Location != "", "location", "must be provided")
	v.Check(len(job.Location) <= 500, "location", "must not be more than 500 bytes long")
	v.Check(job.Status >= 0 && job.Status <= 4, "status", "must be at least 0 and no more than 4")
}