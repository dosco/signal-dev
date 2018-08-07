package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/boltdb/bolt"
)

type ContactTok struct {
	Token      string `json:"token"`
	Relay      string `json:"relay"`
	SupportSMS bool   `json:"supportsSms"`
}

type ContactsReq struct {
	Contacts []string `json:"contacts"`
}

type ContactsResp struct {
	Contacts []ContactTok `json:"contacts"`
}

func directory(w http.ResponseWriter, req *http.Request) {
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

	var cReq ContactsReq
	err = json.Unmarshal(body, &cReq)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	log.Infof("number: %s, contacts: %s", uname, cReq.Contacts)

	var cResp ContactsResp

	err = db.View(func(tx *bolt.Tx) error {
		for _, tok := range cReq.Contacts {
			tokHash, err := base64.RawStdEncoding.DecodeString(tok)
			if err != nil {
				return err
			}
			id := encodeNumber(string(tokHash))
			if tx.Bucket(id) != nil {
				cResp.Contacts = append(cResp.Contacts, ContactTok{Token: tok})
			}
		}
		return nil
	})
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	resp, err := json.Marshal(cResp)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
}
