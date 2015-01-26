package stream

import (
	"github.com/garyburd/go-oauth/oauth"

	"net/http"
	"encoding/json"
	"log"
)

const (
	SAMPLE_URL   = "https://stream.twitter.com/1.1/statuses/sample.json"
	STREAM_URL   = "https://userstream.twitter.com/1.1/user.json"
	USER_URL     = STREAM_URL + "?with=user"

	DETAILS_URL  = "https://api.twitter.com/1.1/account/verify_credentials.json"
)

type auth struct {
	Oauth       *oauth.Client
	Credentials *oauth.Credentials
}

func (a *auth) Details() (string, string, error) {
	req, _ := http.NewRequest("GET", DETAILS_URL, nil)
	req.Header.Set("Authorization", a.Oauth.AuthorizationHeader(a.Credentials, "GET", req.URL, nil))

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return "", "", err
	}

	var c struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return "", "", nil
	}

	return c.Name, c.Url, nil
}

func Auth(consumerKey, consumerSecret, accessToken, accessSecret string) *auth {
	return &auth{
		Oauth: &oauth.Client{
			Credentials: oauth.Credentials{
				Token:  consumerKey,
				Secret: consumerSecret,
			},
		},
		Credentials: &oauth.Credentials{
			Token:  accessToken,
			Secret: accessSecret,
		},
	}
}

type Stream chan Tweet

func Timeline(creds *auth) Stream {
	conn := newConnection(creds)
	err := conn.Open(STREAM_URL)
	if err != nil {
		log.Fatal(err)
	}

	return conn.out
}

func Self(creds *auth) Stream {
	conn := newConnection(creds)
	err := conn.Open(USER_URL)
	if err != nil {
		log.Fatal(err)
	}

	return conn.out
}
