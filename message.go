package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func sendUpdateEvent(me PlayerState) {
	url := "http://localhost:7474/DoAction"

	var jsondat string = fmt.Sprintf(`{"action":{"name":"csgo_data_update"},"args":{"kills":"%d","deaths":"%d","assists":"%d"}}`, me.kills, me.deaths, me.assists)
	var jsonStr2 = []byte(jsondat)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr2))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}

func sendAction(action string) {
	url := "http://localhost:7474/DoAction"

	var dat string = fmt.Sprintf(`{"action":{"name":"%s"}}`, action)
	var jsonStr2 = []byte(dat)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr2))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}
