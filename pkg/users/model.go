package users

import (
	"github.com/rijine/ads-api/pkg/posts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// User full name
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	Url       string `json:"url" bson:"url,omitempty"`
	// used for displaying + auth2 login
	DisplayName     string      `json:"displayName" bson:"displayName,omitempty"`
	ProfileImageUrl string      `json:"profileImageUrl" bson:"profileImageUrl,omitempty"`
	ProfileImage    posts.Image `json:"profileImage" bson:"profileImage,omitempty"`
	Description     string      `json:"description" bson:"description,omitempty"`
	Address         `json:"address" bson:"address,omitempty"`
	Phone           string `json:"phone" bson:"phone,omitempty"`

	Email              string `json:"email" bson:"email,omitempty"`
	Username           string `json:"username" bson:"username,omitempty"`
	Password           string `json:"password" bson:"password,omitempty"`
	VerificationKey    string `json:"verificationKey" bson:"verificationKey,omitempty"`
	VerificationExpiry int64  `json:"verificationExpiry" bson:"verificationExpiry,omitempty"`

	RegisteredOn int64   `json:"registeredOn" bson:"registeredOn,omitempty"`
	IsCompany    bool    `json:"isCompany" bson:"isCompany,omitempty"`
	Rating       float32 `json:"rating" bson:"rating,omitempty"`
	// freq of posting videos
	Frequency   string `json:"frequency" bson:"frequency,omitempty"`
	AvgViews    int64  `json:"avgViews" bson:"avgViews,omitempty"`
	Subscribers string `json:"subscribers" bson:"subscribers,omitempty"`
	// will be used for ML historical
	// Score float32 `json:"score"`
	// Level will be used for trusted customer,
	//their post will go live without verification
	Level int      `json:"level" bson:"level,omitempty"`
	Posts []string `json:"posts" bson:"posts,omitempty"`
}

type Portfolio struct {
	Title string `json:"title" bson:"title,omitempty"`
	// platform = [youtube, facebook, insta]
	Platform string        `json:"platform" bson:"platform,omitempty"`
	Url      string        `json:"url" bson:"url,omitempty"`
	Images   []posts.Image `json:"images" bson:"images,omitempty"` // not sure
}

type PaymentPlan struct {
	Plan      string  `json:"plan" bson:"plan,omitempty"`
	UpdatedOn int32   `json:"updatedOn" bson:"updatedOn,omitempty"`
	Expiry    int64   `json:"expiry" bson:"expiry,omitempty"`
	Pay       float32 `json:"pay" bson:"pay,omitempty"`
	NoOfPosts int32   `json:"noOfPosts" bson:"noOfPosts,omitempty"`
}

type Address struct {
	HouseNo  string `json:"houseNo" bson:"houseNo,omitempty"`
	Street   string `json:"street" bson:"street,omitempty"`
	City     string `json:"city" bson:"city,omitempty"`
	District string `json:"district" bson:"district,omitempty"`
	State    string `json:"state" bson:"state,omitempty"`
	Country  string `json:"country" bson:"country,omitempty"`
	PinCode  string `json:"pinCode" bson:"pinCode,omitempty"`
}
