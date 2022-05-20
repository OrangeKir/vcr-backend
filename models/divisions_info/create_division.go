package divisions_info

type CreateDivisionRequest struct {
	DivisionCode       int
	OrganisationName   string
	DivisionName       string
	HeadDivisionCode   *int
	DivisionSupervisor string

	Token string
}
