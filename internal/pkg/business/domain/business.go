package business

import "github.com/google/uuid"

type Business struct {
	Id       uuid.UUID
	Name     string
	Banner   string
	Image    string
	State    string
	HourHand string
	About    string
	Location string
}
