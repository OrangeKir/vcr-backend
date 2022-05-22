package research_works_info

import "time"

type CreateResearchWorkRequest struct {
	Code                     string
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

	Token string
}
