package model

type Address struct {
	HouseNo  string `json:"houseNo"`
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	State    string `json:"state"`
	Country  string `json:"country"`
	PinCode  string `json:"pinCode"`
}
