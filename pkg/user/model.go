package user

import "github.com/rijine/ads-api/pkg/post"

type User struct {
	Id string `json:"_id"`
	// User full name
	Name string `json:"name"`
	Url  string `json:"url"`
	// used for displaying + auth2 login
	DisplayName     string     `json:"displayName"`
	ProfileImageUrl string     `json:"profileImageUrl"`
	ProfileImage    post.Image `json:"profileImage"`
	Description     string     `json:"description"`
	Address         `json:"address"`
	Phone           string `json:"phone"`

	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	VerificationKey    string `json:"verificationKey"`
	VerificationExpiry string `json:"verificationExpiry"`

	RegisteredOn string  `json:"registeredOn"`
	IsCompany    string  `json:"isCompany"`
	Rating       float32 `json:"rating"`
	// freq of posting videos
	Frequency   string `json:"frequency"`
	AvgViews    int32  `json:"avgViews"`
	Subscribers string `json:"subscribers"`
	// will be used for ML historical
	// Score float32 `json:"score"`
	// Level will be used for trusted customer,
	//their post will go live without verification
	Level int      `json:"level"`
	Posts []string `json:"posts"`
}

type Portfolio struct {
	Title string `json:"title"`
	// platform = [youtube, facebook, insta]
	Platform string       `json:"platform"`
	Url      string       `json:"url"`
	Images   []post.Image `json:"images"` // not sure
}

type PaymentPlan struct {
	Plan      string  `json:"plan"`
	UpdatedOn int32   `json:"updatedOn"`
	Expiry    int32   `json:"expiry"`
	Pay       float32 `json:"pay"`
	Posts     int32   `json:"posts"`
}

type Address struct {
	HouseNo  string `json:"houseNo"`
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	Country  string `json:"country"`
	PinCode  string `json:"pinCode"`
}
