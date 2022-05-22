package publications_info

type UpdatePublicationRequest struct {
	PublicationId     int
	PublishingHouseId int
	TopicName         string
	AuthorsLogins     []string
}
