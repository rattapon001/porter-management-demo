package domain

import "time"

type JobId int64
type JobStatus int8

const (
	Pending   JobStatus = 1
	Accepted  JobStatus = 2
	Working   JobStatus = 3
	Completed JobStatus = 4
	Cancel    JobStatus = -1
)

type Location struct {
	From string
	To   string
}

type Patient struct {
	Name string
	HN   string
}

type Porter struct {
	Code string
	Name string
}

type Job struct {
	Id       JobId
	Version  int
	Status   JobStatus
	Accepted bool
	Location Location
	Patient  Patient
	Porter   Porter
	CheckIn  time.Time
	CheckOut time.Time
}

func CreateNewJob(location Location, patient Patient) *Job {
	return &Job{
		Status:   Pending,
		Location: location,
		Patient:  patient,
	}
}

func AcceptJob(job *Job, porter Porter) {
	job.Status = Accepted
	job.Porter = porter
	job.Accepted = true
}
