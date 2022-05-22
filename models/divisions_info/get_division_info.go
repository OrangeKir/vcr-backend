package divisions_info

type GetDivisionInfoRequest struct {
	DivisionCode int

	Token string
}

type GetDivisionInfoResponse struct {
	OrganisationName   string
	DivisionName       string
	HeadDivisionCode   int
	HeadDivisionName   string
	DivisionSupervisor string
}
