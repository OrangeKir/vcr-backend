package persons_info

type GetPersonsResponse struct {
	PersonsInfo []PersonMainInfo
}

type PersonMainInfo struct {
	Login      string
	FirstName  string
	SecondName string
	ThirdName  string
}
