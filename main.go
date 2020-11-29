package main

import (
	"bytes"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	log.Info("hello world")

	http.HandleFunc("/", EmailSignIn)
	_ = http.ListenAndServe(":9001", nil)
}

func EmailSignIn(w http.ResponseWriter, r *http.Request) {
	googleApiKey := "myapikey"
	urlString := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword/?key=" + googleApiKey

	u, err := url.Parse(r.URL.String())

	jsonString := "{'email':'" + u.Query().Get("email") + "', 'password': '" + u.Query().Get("password") + "', 'returnSecureToken': true}"

	// Make API Call
	req, err := http.NewRequest("POST", urlString, bytes.NewBufferString(jsonString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Info("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(body))

	w.WriteHeader(resp.StatusCode)
	w.Write(body)

}
