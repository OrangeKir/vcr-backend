package conferences_info

type GetConferencesResponse struct {
	Conferences []Conference
}

type Conference struct {
	Id   int
	Name string
}
