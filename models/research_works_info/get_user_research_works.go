package research_works_info

type GetUserResearchWorksRequest struct {
	Login string

	Token string
}

type GetUserResearchWorksResponse struct {
	ResearchWorksMainInfo []ResearchWorkMainInfo
}

type ResearchWorkMainInfo struct {
	Code           string
	ShortWorkTitle string
}
