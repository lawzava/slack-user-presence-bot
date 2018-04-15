package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type userRawData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Presence string `json:"presence"`
}

type usersRawResponse struct {
	Members []userRawData `json:"members"`
}

func checkUsersPresence(token string) ([]userRawData, error) {
	slackUserListURL := "https://slack.com/api/users.list?presence=true&token=" + token
	res, err := http.Get(slackUserListURL)
	if err != nil {
		return nil, fmt.Errorf("error while doing http request: %v", err)
	}

	users, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error while reading http response: %v", err)
	}

	var usersData usersRawResponse
	err = json.Unmarshal(users, &usersData)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling http response: %v", err)
	}

	return usersData.Members, nil
}
