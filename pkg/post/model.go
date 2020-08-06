package post

type Post struct {
	Id          string   `json:"id"`
	UserId      string   `json:"userId"`
	Url         string   `json:"url"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Images      []Image  `json:"images"`
	Price       float32  `json:"price"`
	Tags        []string `json:"tags"`

	// youtube, facebook, instagram
	Platforms     []string `json:"platforms"`
	Language      string   `json:"language"`
	ExpectedViews int64    `json:"expectedViews"`
	Duration      int      `json:"duration"`

	Status string `json:"status"`
	// Can be remove by adding one more status type
	IsActive       bool  `json:"isActive"`
	AvailableFrom  int64 `json:"availableFrom"`
	PublishingDate int64 `json:"publishingDate"`
	CreatedAt      int64 `json:"createdAt"`

	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Image struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumbUrl"`
	Description string `json:"description"`
}
