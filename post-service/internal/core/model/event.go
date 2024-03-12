package model

type PostEvent struct {
	Title         string
	PostId        int
	SubscriberIds []int
}
