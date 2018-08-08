package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Key struct {
	KeyID     int64  `json:"keyId"`
	PublicKey string `json:"publicKey"`
}

type SignedKey struct {
	KeyID     int64  `json:"keyId"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

type PreKeys struct {
	SignedPreKey  SignedKey `json:"signedKey"`
	IdentityKey   string    `json:"identityKey`
	LastResortKey Key       `json:"lastResortKey"`
	PreKeys       []Key     `json:"preKeys"`
}

func registerKeys(w http.ResponseWriter, req *http.Request) {
	uname, _, ok := req.BasicAuth()
	if !ok {
		http.Error(w, "auth fail", 401)
		return
	}

	log.Infof("number: %s", uname)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	var preKeys PreKeys
	err = json.Unmarshal(body, &preKeys)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}
	log.Infof("%+v\n", preKeys)
	log.Infof("%s\n", body)

	id := encodeNumber(uname)
	writeDB(id[:], []byte("k"), body)
	w.WriteHeader(200)
}

func recipientKeys(w http.ResponseWriter, req *http.Request) {
	uname, _, ok := req.BasicAuth()
	if !ok {
		http.Error(w, "auth fail", 401)
		return
	}

	log.Infof("number: %s", uname)

	id := encodeNumber(uname)
	value, err := readDB(id[:], []byte("k"))
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	var preKeys PreKeys
	err = json.Unmarshal(value, &preKeys)
	if err != nil {
		http.Error(w, "server error", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(value)
}
