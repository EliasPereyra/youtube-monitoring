package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Kind			string			`json:"kind"`
	Items			[]Items			`json:"items"`
}

type Stats struct {
	ViewsCount				int		`json:"views"`
	SuscribersCount		int		`json:"suscribers"`
	VideosCount				int		`json:"videos"`
}

type Items struct {
	Kind	string					`json:"kind"`
	Etag	string					`json:"etag"`
	Id		string					`json:"id"`
}

func GetSuscribers() (Items, error) {
	var YOUTUBE_API_URL = "https://www.googleapis.com/youtube/v3/channels"
	var response Response

	req, err := http.NewRequest("GET", YOUTUBE_API_URL, nil)

	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	q := req.URL.Query()

	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return Items{}, err
	}

	defer res.Body.Close()

	fmt.Println("Response Status", res.Status)

	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return Items{}, err
	}

	return response.Items[0], nil
}