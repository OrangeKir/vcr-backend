package research_works_info

type WorkType string

const (
	ResearchWork        WorkType = "ResearchWork"
	ResearchWorkPart    WorkType = "ResearchWorkPart"
	DevelopmentWork     WorkType = "DevelopmentWork"
	DevelopmentWorkPart WorkType = "DevelopmentWorkPart"
	Service             WorkType = "Service"
	Grant               WorkType = "Grant"
	Other               WorkType = "Other"
)
