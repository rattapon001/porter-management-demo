package domain

import (
	"time"

	"github.com/google/uuid"
)

type JobId string
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
	Id       JobId     `bson:"_id"`
	Version  int       `bson:"version"`
	Status   JobStatus `bson:"status"`
	Accepted bool      `bson:"accepted"`
	Location Location  `bson:"location"`
	Patient  Patient   `bson:"patient"`
	Porter   Porter    `bson:"porter"`
	CheckIn  time.Time `bson:"check_in"`
	CheckOut time.Time `bson:"check_out"`
}

func CreateNewJob(location Location, patient Patient) *Job {
	return &Job{
		Id:       JobId(uuid.New().String()),
		Status:   Pending,
		Location: location,
		Patient:  patient,
		Version:  1,
	}
}

func AcceptJob(job *Job, porter Porter) {
	job.Status = Accepted
	job.Porter = porter
	job.Accepted = true
}
