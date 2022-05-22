package staff_info

import "time"

type GetStaffInfoRequest struct {
	PersonLogin string

	Token string
}

type GetStaffInfoResponse struct {
	StaffRecords []StaffRecord
}

type StaffRecord struct {
	Position                       string
	WorkingRate                    float32
	ProfessionalQualificationGroup string
	QualificationLevel             string
	PositionType                   PositionType
	OrganisationName               string
	DivisionId                     int
	DivisionName                   string
	SupervisorLogin                *string
	SupervisorFirstName            string
	SupervisorSecondName           string
	SupervisorThirdName            string
	IsSupervisor                   string
	ContractEndDate                time.Time
	Status                         StatusType
	StatusChangeReason             StatusChangeReasonType
}
