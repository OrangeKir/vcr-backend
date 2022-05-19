package persons_info

import "time"

type CreatePersonRequest struct {
	Login          string
	FirstName      string
	SecondName     string
	ThirdName      string
	DateOfBirth    time.Time
	Education      string
	AcademicDegree string
	AcademicRank   string
	ContactNumber  string
	ContactEmail   string

	Token string
}
