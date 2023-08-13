package types

type Message struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	User    string `json:"user"`
}

type MessageCreateDto struct {
	Message string `json:"message"`
	UserId  string `json:"-"`
}
