package publications_info

type CreatePublicationRequest struct {
	PublishingHouseId int
	TopicName         string
	AuthorsLogins     []string
}
