package domain

type PorterId int64
type PorterStatus int8

type Porter struct {
	Id         PorterId
	PorterCode string
	Status     PorterStatus
	Name       string
}

const (
	Available PorterStatus = 1
	Working   PorterStatus = 2
	Offline   PorterStatus = 3
)

func NewPorter(
	name string,
) (Porter, error) {
	return Porter{
		Name:   name,
		Status: Offline,
	}, nil
}
