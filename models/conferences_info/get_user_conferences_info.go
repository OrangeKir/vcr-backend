package conferences_info

type GetUserConferencesInfoRequest struct {
	Login string
}

type GetUserConferencesInfoResponse struct {
	Conferences []UserConference
}

type UserConference struct {
	Id         int
	Name       string
	TopicNames []string
}
