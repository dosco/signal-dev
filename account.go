package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func createAccount(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Infof("transport: %s, number: %s", vars["sms"], vars["number"])

	w.WriteHeader(200)
}

func verifyCode(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Infof("verification_code: %s", vars["verification_code"])

	/*{
	  "signalingKey" : "{base64_encoded_52_byte_key}",
	  "supportsSms" : false,
	  "registrationId" : "{14-bit number}"
		}*/
	w.WriteHeader(200)
}

type Attributes struct {
	AuthKey         string `json:"AuthKey"`
	FetchesMessages bool   `json:"fetchesMessages"`
	Voice           bool   `json:"voice"`
	Video           bool   `json:"video"`
	RegistrationID  string `json:"registrationId"`
	SignalingKey    string `json:"signalingKey"`
}

func saveAttributes(w http.ResponseWriter, req *http.Request) {
	uname, _, ok := req.BasicAuth()
	if !ok {
		http.Error(w, "auth fail", 401)
		return
	}

	vars := mux.Vars(req)
	log.Infof("number: %s, device_id: %s", vars["number"], vars["device_id"])

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	var attrib Attributes
	err = json.Unmarshal(body, &attrib)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	id := encodeNumber(uname)
	writeDB(id[:], []byte("a"), body)
	w.WriteHeader(200)
}
