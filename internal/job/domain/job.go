package domain

import (
	"time"

	"github.com/google/uuid"
)

type JobId string
type JobStatus int8
type JobEvent string

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

const (
	JobCreated   JobEvent = "jobCreated"
	JobAccepted  JobEvent = "jobAccepted"
	JobStarted   JobEvent = "jobStarted"
	JobCompleted JobEvent = "jobCompleted"
)

type EventDataCreated struct {
	ID       JobId     `bson:"id"`
	Location Location  `bson:"location"`
	Patient  Patient   `bson:"patient"`
	Version  int       `bson:"version"`
	Name     string    `bson:"name"`
	Porter   Porter    `bson:"porter"`
	CheckIn  time.Time `bson:"check_in"`
	CheckOut time.Time `bson:"check_out"`
}

type Event struct {
	Type JobEvent         `bson:"type"`
	Data EventDataCreated `bson:"data"`
}

type Job struct {
	ID       JobId     `bson:"_id,omitempty"`
	Version  int       `bson:"version"`
	Status   JobStatus `bson:"status"`
	Accepted bool      `bson:"accepted"`
	Location Location  `bson:"location"`
	Patient  Patient   `bson:"patient"`
	Porter   Porter    `bson:"porter"`
	CheckIn  time.Time `bson:"check_in"`
	CheckOut time.Time `bson:"check_out"`
	Event    []Event   `bson:"event"`
}

func CreateNewJob(location Location, patient Patient) *Job {
	return &Job{
		ID:       JobId(uuid.New().String()),
		Status:   Pending,
		Location: location,
		Patient:  patient,
		Version:  0,
	}
}

func AcceptJob(job *Job, porter Porter) {
	job.Status = Accepted
	job.Porter = porter
	job.Accepted = true
}

func (j *Job) AppendEvent(event Event) {
	j.Event = append(j.Event, event)
}
