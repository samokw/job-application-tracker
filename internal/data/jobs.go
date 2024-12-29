package data

import (
	"time"
)

type Status int

const (
	Saved Status = iota
	Applied
	Interview
	Accepted
	Rejected
)

type Job struct {
	ID              int64     `json:"id"`
	CompanyName     string    `json:"company_name"`
	Position        string    `json:"position"`
	Description     string    `json:"description,omitempty"`
	ApplicationDate time.Time `json:"appliaction_date"`
	Location        string    `json:"location"`
	Link            string    `json:"link,omitempty"`
	Status          Status    `json:"status"`
}
