package staff_info

import "time"

type UpdateStaffInfoRequest struct {
	Id                             int
	Position                       string
	WorkingRate                    float32
	ProfessionalQualificationGroup string
	QualificationLevel             string
	PositionType                   PositionType
	OrganisationName               string
	DivisionId                     int
	SupervisorLogin                string
	IsSupervisor                   string
	ContractEndDate                time.Time
	Status                         StatusType
	StatusChangeReason             StatusChangeReasonType

	Token string
}
