package staff_info

type PositionType string
type StatusType string
type StatusChangeReasonType string

const (
	Main                PositionType = "Main"
	InternalCombination PositionType = "InternalCombination"
	ExternalCombination PositionType = "ExternalCombination"
)

const (
	Active   StatusType = "Active"
	Inactive StatusType = "Inactive"
)

const (
	Hired    StatusChangeReasonType = "Hired"
	Fired    StatusChangeReasonType = "Fired"
	Transfer StatusChangeReasonType = "Transfer"
)
