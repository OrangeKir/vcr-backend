package research_works_info

import "time"

type GetResearchWorkInfoRequest struct {
	Code int

	Token string
}

type GetResearchWorkInfoResponse struct {
	LeadCustomer             string
	Customer                 string
	FullWorkTitle            string
	ShortWorkTitle           string
	ContractNumber           int
	ContractDate             time.Time
	TopicNumber              int
	WorkType                 WorkType
	WorkPrice                int
	StartDate                time.Time
	EndDate                  time.Time
	StagesNumber             int
	WorkPriceStructureId     int
	ResponsibleExecutorLogin string
}
