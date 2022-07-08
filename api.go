package rubbrband_go

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type User struct {
	Name       string
	ProfileUrl string
}

type Entry struct {
	Value string
	User  User
}

// var FIAT_ENDPOINT = "https://api.rubbrband.com/"
var FIAT_ENDPOINT = "http://localhost:4200/"

var ApiKey string

func SetApiKey(Key string) {
	ApiKey = Key
}

func Store(key string, value string, user User) string {
	if ApiKey == "" {
		return "No API Key set! Call SetApiKey with your Rubbrband API Key."
	}

	entry := Entry{
		Value: value,
		User:  user,
	}
	body, _ := json.Marshal(entry)

	resp, err := http.Post(FIAT_ENDPOINT+"put/"+key+"?api_key="+ApiKey, "application/json", bytes.NewBuffer(body))
	if resp.Status != "200 OK" {
		return "Unsuccessful, please check your API key and user info"
	}
	if err != nil {
		panic(err)
	}

	return value
}

func Replay(key string) string {
	if ApiKey == "" {
		return "No API Key set! Call SetApiKey with your Rubbrband API Key."
	}

	resp, err := http.Get(FIAT_ENDPOINT + "get/" + key + "?api_key=" + ApiKey)
	if resp.Status != "200 OK" {
		return "Unsuccessful, please check your API key, and cache key"
	}

	if err != nil {
		panic(err)
	}

	var data Entry
	json.NewDecoder(resp.Body).Decode(&data)

	return data.Value
}

func Delete(key string) bool {
	if ApiKey == "" {
		return false
	}

	resp, err := http.Post(FIAT_ENDPOINT+"delete/"+key+"?api_key="+ApiKey, "application/json", nil)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}
