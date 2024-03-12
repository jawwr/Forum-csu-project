package model

type PostEvent struct {
	Title         string
	PostId        int
	SubscriberIds []int
}

type EventResponse struct {
	Id     int
	Title  string
	PostId int
}
