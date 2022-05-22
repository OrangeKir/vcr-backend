package publications_info

type GetUserPublicationsInfoByIdRequest struct {
	Id int
}

type GetUserPublicationsInfoByIdResponse struct {
	PublicationId       int
	TopicName           string
	PublishingHouseId   int
	PublishingHouseName string
	AuthorsLogins       []string
}
