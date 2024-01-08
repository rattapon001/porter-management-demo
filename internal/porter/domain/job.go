package domain

type JobId string
type JobStatus int8
type JobEvent string

type Location struct {
	From string
	To   string
}

type Patient struct {
	Name string
	HN   string
}

const (
	Accepted  JobStatus = 1
	Completed JobStatus = 2
	Cancel    JobStatus = -1
)

const (
	JobCreated   JobEvent = "jobCreated"
	JobAccepted  JobEvent = "jobAccepted"
	JobStarted   JobEvent = "jobStarted"
	JobCompleted JobEvent = "jobCompleted"
)

type EventDataCreated struct {
	ID       JobId    `bson:"id"`
	Location Location `bson:"location"`
	Patient  Patient  `bson:"patient"`
	Version  int      `bson:"version"`
	Name     string   `bson:"name"`
	Porter   Porter   `bson:"porter"`
}

type Event struct {
	Type JobEvent         `bson:"type"`
	Data EventDataCreated `bson:"data"`
}

type Job struct {
	ID       JobId     `bson:"_id,omitempty"`
	SourceID JobId     `bson:"source_id"`
	Version  int       `bson:"version"`
	Status   JobStatus `bson:"status"`
	Location Location  `bson:"location"`
	Patient  Patient   `bson:"patient"`
	Porter   Porter    `bson:"porter"`
	Name     string    `bson:"name"`
	Event    []Event   `bson:"event"`
}
