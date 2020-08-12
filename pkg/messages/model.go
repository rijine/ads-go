package messages

import "github.com/99designs/gqlgen/graphql"

type Enquiry struct {
	Id     string `json:"id"`
	PostId string `json:"postId"`
	// populate userid
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Subject    string `json:"subject"`
	Content    string `json:"content"`
	MessagedAt int64  `json:"messagedAt"`
	IsRead     bool   `json:"isRead"`
}

type Chat struct {
	Id         string    `json:"id"`
	UserId     string    `json:"userId"`
	ReceiverId string    `json:"receiverId"`
	Messages   []Message `json:"messages"`
}

type Message struct {
	Content    string         `json:"content"`
	Image      graphql.Upload `json:"image"`
	IsRead     bool           `json:"isRead"`
	MessagedAt int64          `json:"messageAt"`
}

type Image struct {
	Url      string `json:"url"`
	ThumbUrl string `json:"thumbUrl"`
}
