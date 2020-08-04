package enquiry

type Message struct {
	Id         string `json:"id"`
	PostId     string `json:"postId"`
	UserId     string `json:"userId"`
	Subject    string `json:"subject"`
	Content    string `json:"content"`
	MessagedAt string `json:"messageAt"`

	IsExternal bool   `json:"isExternal"`
	SenderId   string `json:"senderId"`
	Name       string `json:"Name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
