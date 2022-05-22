package publications_info

type GetUserPublicationsInfoByLoginRequest struct {
	Login string
}

type GetUserPublicationsInfoByLoginResponse struct {
	Publications []Publication
}

type Publication struct {
	PublicationId       int
	TopicName           string
	PublishingHouseName string
}
