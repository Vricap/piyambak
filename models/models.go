package models

type Message struct {
	ID        string `josn:"id"`
	Text      string `json:"text"`
	User      string `json:"user"`
	CreatedAt string `json:"createdAt"`
}
