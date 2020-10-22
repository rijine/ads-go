package youtubeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Service interface {
	GetSubscribers(userId string) (Items, error)
}
type service struct{}

func NewYoutubeApiService() Service {
	fmt.Print("youtubeapi service")
	return &service{}
}

func (s *service) GetSubscribers(userId string) (Items, error) {
	var response Response
	// We want to craft a new GET request that will include the query parameters we want
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	// here we define the query parameters and their respective values
	q := req.URL.Query()
	// notice how I'm using os.Getenv() to pick up the environment
	// variables that we defined earlier. No hard coded credentials here
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("forUsername", userId)
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	// finally we make the request to the URL that we have just
	// constructed
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Items{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	// we then read in all of the body of the
	// JSON response
	body, _ := ioutil.ReadAll(resp.Body)
	// and finally unmarshal it into an Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Items{}, err
	}
	// we only care about the first Item in our
	// Items array, so we just send that back
	return response.Items[0], nil
}
