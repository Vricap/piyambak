package models

type Room struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	User      string `json:"user"`
	RoomId    string `json:"roomId"` // hmmm
	CreatedAt string `json:"createdAt"`
}
