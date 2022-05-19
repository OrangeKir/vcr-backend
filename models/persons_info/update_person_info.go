package persons_info

type UpdatePersonInfoRequest struct {
	Login          string
	Education      string
	AcademicDegree string
	AcademicRank   string
	ContactNumber  string
	ContactEmail   string

	Token string
}
