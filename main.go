package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var googleApiKey string


func main() {
	// Add your Google API key
	googleApiKey = os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		panic("need the google api key")
	}

	log.Printf("ready to accept requests on port 9001")

	http.HandleFunc("/", EmailSignIn)
	_ = http.ListenAndServe(":9001", nil)
}

func EmailSignIn(w http.ResponseWriter, r *http.Request) {
	urlString := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword/?key=" + googleApiKey

	// Do we have our mandatory params?
	u, err := url.Parse(r.URL.String())
	if u.Query().Get("email") == "" {
		log.Println("email is missing, request failed")
		w.WriteHeader(400)
		w.Write([]byte("'email' missing from request"))
		return
	}
	if u.Query().Get("password") == "" {
		log.Println("password is missing, request failed")
		w.WriteHeader(400)
		w.Write([]byte("'password' missing from request"))
		return
	}

	// Construct the upstream query and make API Call
	jsonString := "{'email':'" + u.Query().Get("email") + "', 'password': '" + u.Query().Get("password") + "', 'returnSecureToken': true}"
	req, err := http.NewRequest("POST", urlString, bytes.NewBufferString(jsonString))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	// Will write the status code from Google API to proxy
	// This is important because a 200 will allow the user to sign in
	// and a non-200 will tell the downstream that the authentication failed
	// Will echo the response from Google API to downstream
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
