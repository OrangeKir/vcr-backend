package divisions_info

type UpdateDivisionInfoRequest struct {
	DivisionCode       int
	OrganisationName   string
	DivisionName       string
	HeadDivisionCode   int
	DivisionSupervisor string
}
