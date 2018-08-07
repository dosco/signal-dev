package main

import (
	"crypto/sha1"
	"os"
	// "fmt"
	// "io"

	"net/http"
	"net/http/httputil"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	logging "github.com/op/go-logging"
)

//https://github.com/signalapp/Signal-Server/wiki/API-Protocol

var (
	db  *bolt.DB
	log *logging.Logger
)

func encodeNumber(e164 string) []byte {
	b := sha1.Sum([]byte(e164))
	return b[0:10]
}

func notFound(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Errorf("%s", err)
	} else {
		log.Info("%q", dump)
	}
}

func main() {
	var err error

	log = logging.MustGetLogger("signal-dev")
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Formatter)

	db, err = bolt.Open("dev.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc(
		"/v1/accounts/{transport}/code/{number}",
		createAccount).Methods("GET")

	r.HandleFunc("/v1/acccounts/code/{verification_code}",
		verifyCode).Methods("PUT")

	r.HandleFunc("/v2/keys",
		registerKeys).Methods("PUT")

	r.HandleFunc("/v1/accounts/code/{code}",
		saveAttributes).Methods("PUT")

	r.HandleFunc("/v1/accounts/attributes/",
		saveAttributes).Methods("PUT")

	r.HandleFunc("/v1/messages/{number}",
		submitMessages).Methods("PUT")

	r.HandleFunc("/v1/keys/{number}/{device_id}",
		recipientKeys).Methods("GET")

	r.HandleFunc("/v1/directory/tokens",
		directory).Methods("PUT")

	r.NotFoundHandler = http.HandlerFunc(notFound)
	/*
		//w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			},
		}
		srv := &http.Server{
			Addr:         ":4443",
			Handler:      r,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	*/

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Info("Signal-DEV listening on :8080")
	log.Fatal(srv.ListenAndServe())

}
