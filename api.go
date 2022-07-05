package rubbrband_go

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func Store(key string, value string, local bool, user User) string {
	if ApiKey == "" {
		return "No API Key set! Call SetApiKey with your Rubbrband API Key."
	}

	entry := Entry{
		Value: value,
		User:  user,
	}
	body, _ := json.Marshal(entry)

	_, err := http.Post(FIAT_ENDPOINT+"put/"+key+"?api_key="+ApiKey, "application/json", bytes.NewBuffer(body))

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

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//Failed to read response.
		panic(err)
	}

	//Convert bytes to String and print
	jsonStr := string(body)

	return jsonStr
}

func Delete(key string) bool {
	if ApiKey == "" {
		return false
	}

	resp, err := http.Post(FIAT_ENDPOINT+"delete/"+key+"?api_key="+ApiKey, "application/json", nil)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusCreated {
		return true
	} else {
		return false
	}
}
