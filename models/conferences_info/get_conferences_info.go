package conferences_info

type GetUserConferencesInfoRequest struct {
	Login string
}

type GetUserConferencesInfoResponse struct {
	Conferences []Conference
}

type Conference struct {
	Name       string
	TopicNames []string
}
