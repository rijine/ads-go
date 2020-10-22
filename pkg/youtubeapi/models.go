package youtubeapi

// To capture main response
// store items
type Response struct {
	Kind  string  `json:"kind" bson:"kind,omitempty"`
	Items []Items `json:"items" bson:"items,omitempty"`
}

// Items stores the ID + Statistics for
// a given channel
type Items struct {
	Kind  string `json:"kind" bson:"kind,omitempty"`
	ID    string `json:"id" bson:"id,omitempty"`
	Stats Stats  `json:"statistics" bson:"statistics,omitempty"`
}

// Stats stores the information we care about
// so how many views the channel has, how many subscribers
// how many video etc.
type Stats struct {
	Views       string `json:"viewCount" bson:"viewCount,omitempty"`
	Subscribers string `json:"subscriberCount" bson:"subscriberCount,omitempty"`
	Videos      string `json:"videoCount" bson:"videoCount,omitempty"`
}
