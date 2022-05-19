package persons_info

import (
	"time"
)

type GetPersonInfoRequest struct {
	Login string
}

type GetPersonInfoResponse struct {
	FirstName      string
	SecondName     string
	ThirdName      string
	DateOfBirth    time.Time
	Education      string
	AcademicDegree string
	AcademicRank   string
	ContactNumber  string
	ContactEmail   string
}
