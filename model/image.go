package model

type Image struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumbUrl"`
	Description string `json:"description"`
}
