package publications_info

type GetPublishingHousesInfoResponse struct {
	PublishingHousesInfo []PublishingHouseInfo
}

type PublishingHouseInfo struct {
	Id   int
	Name string
}
