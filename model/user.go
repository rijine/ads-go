package model

type User struct {
	Id string `json:"_id"`
	// User full name
	Name string `json:"name"`
	// used for displaying + auth2 login
	DisplayName  string `json:"displayName"`
	ProfileImage string `json:"profileImage"`
	Description  string `json:"description"`
	Address      `json:"address"`
	Phone        string `json:"phone"`

	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	VerificationKey    string `json:"verificationKey"`
	VerificationExpiry string `json:"verificationExpiry"`

	RegisteredOn string  `json:"registeredOn"`
	IsCompany    string  `json:"isCompany"`
	Rating       float32 `json:"rating"`
	// ???
	AvgViews int32 `json:"avgViews"`
	// will be used for ??
	Score float32 `json:"score"`
	// Level will be used for trusted customer,
	//their post will go live without verification
	Level string `json:"level"`
}

type Portfolio struct {
	Name string `json:"name"`
	// type = [websiteUrl, videoUrl, images]
	Type   string  `json:"type"`
	Url    string  `json:"url"`
	Images []Image `json:"images"` // not sure
}

type PaymentPlan struct {
	Plan      string  `json:"plan"`
	UpdatedOn int32   `json:"updatedOn"`
	Expiry    int32   `json:"expiry"`
	Pay       float32 `json:"pay"`
}
