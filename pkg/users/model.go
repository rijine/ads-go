package users

import (
	"github.com/rijine/ads-api/pkg/posts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	// User full name
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Url       string `json:"url"`
	// used for displaying + auth2 login
	DisplayName     string      `json:"displayName"`
	ProfileImageUrl string      `json:"profileImageUrl"`
	ProfileImage    posts.Image `json:"profileImage"`
	Description     string      `json:"description"`
	Address         `json:"address"`
	Phone           string `json:"phone"`

	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	VerificationKey    string `json:"verificationKey"`
	VerificationExpiry int64  `json:"verificationExpiry"`

	RegisteredOn int64   `json:"registeredOn"`
	IsCompany    bool    `json:"isCompany"`
	Rating       float32 `json:"rating"`
	// freq of posting videos
	Frequency   string `json:"frequency"`
	AvgViews    int64  `json:"avgViews"`
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
	Platform string        `json:"platform"`
	Url      string        `json:"url"`
	Images   []posts.Image `json:"images"` // not sure
}

type PaymentPlan struct {
	Plan      string  `json:"plan"`
	UpdatedOn int32   `json:"updatedOn"`
	Expiry    int64   `json:"expiry"`
	Pay       float32 `json:"pay"`
	NoOfPosts int32   `json:"noOfPosts"`
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
