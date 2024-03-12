package model

type HandlePost struct {
	Id    int        `json:"id"`
	Title string     `json:"title"`
	Text  string     `json:"text"`
	User  HandleUser `json:"user"`
}
