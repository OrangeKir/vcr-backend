package publications_info

type CreateUsersPublicationRequest struct {
	PublishingHouseId int
	TopicName         string
	AuthorsLogins     []string

	Token string
}
