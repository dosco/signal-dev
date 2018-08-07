package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Type                      string `json:"type"`
	DestinationDeviceId       int64  `json:"destinationDeviceId"`
	DestinationRegistrationId int64  `json:"destinationRegistrationId"`
	Body                      string `json:"body"`
	timestamp                 string `json:"timestamp"`
}

type Messages struct {
	Relay    string    `json:"relay"`
	Messages []Message `json:"messages"`
}

func submitMessages(w http.ResponseWriter, req *http.Request) {
	uname, _, ok := req.BasicAuth()
	if !ok {
		http.Error(w, "auth fail", 401)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	var messages Messages
	err = json.Unmarshal(body, &messages)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	log.Infof("number: %s, messages_count: %d", uname, len(messages.Messages))

	id := encodeNumber(uname)
	writeDB(id[:], []byte("m"), body)
	w.WriteHeader(200)
}
