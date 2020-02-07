package handlers

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"slack-bot/config"
)

func GetEnvironments(w http.ResponseWriter, r *http.Request) {
	botConfig := config.GetConfig()

	apiKey := botConfig.CycleAPIKey
	hubID := botConfig.HubID

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.cycle.io/v1/environments", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("X-Hub-Id", hubID)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	w.Write(body)
}
