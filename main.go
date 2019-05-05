package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const api_server = "api.token.wtf"
const api_key = "5d5eb8f166fb2c0b"
const api_secret = "30878c87d4e7f96e98dd4092"
const api_token = "93f4f193-62a5-4bec-acaa-f17322bdc65e"

func main() {

	// hash := sha256.New()
	// hash.Write([]byte("28813502-6a08-40bd-90a5-79c22691852d"))
	// fmt.Printf("https://token.wtf/stages/4-%s.html\n\n", hex.EncodeToString(hash.Sum(nil)))

	// response, err := getScore()
	// fmt.Printf("Resp: %s\nError: %s\n", response, err)

	// response, err = validateToken(api_token)
	// fmt.Printf("Resp: %s\nError: %s\n", response, err)

	// Stage 1:
	// response, err := claimToken("93f4f193-62a5-4bec-acaa-f17322bdc65e")
	// fmt.Printf("Resp: %s\nError: %s\n", response, err)

	// Stage 2:
	// https://token.wtf/stages/1-08f12daef7c0b48a4530c21fbfd713db8ae6cfda8e982799bd775adf76eeb636.html
	// response, err := claimToken("a6ad2799-8a51-4ddb-bb55-83504106106e")
	// fmt.Printf("Resp: %s\nError: %s\n", response, err)

	// Stage 3:
	// https://token.wtf/stages/2-1cb88db26fd2d30eb312d9ddd732e2c6fc58373888279e40932b2deffec4a162.html
	//response, err := claimToken("365ffb4c-4ef5-a030-5fa9-740f1518c336")
	//fmt.Printf("Resp: %s\nError: %s\n", response, err)

	// Stage 4:
	// https://token.wtf/stages/3-eccbfef060736326603848438ab4bc5e5c542aadff4d8495e005ee2d0ea09d75.html
	// response, err := claimToken("28813502-6a08-40bd-90a5-79c22691852d")
	// fmt.Printf("Resp: %s\nError: %s\n", response, err)

	// Stage 5:
	//
	response, err := claimToken("35251dc7-544b-4267-a91b-3db0c8842113")
	fmt.Printf("Resp: %s\nError: %s\n", response, err)

}

func hmacSignature(url, body string) (string) {

	api_request := fmt.Sprintf("%s\n%s\n%s\n%s\n", api_server, url, api_key, body)

	api_signature := hmac.New(sha256.New, []byte(api_secret))
	api_signature.Write([]byte(fmt.Sprint(api_request)))

	return hex.EncodeToString(api_signature.Sum(nil))
}

func validateToken(token string) (string, error) {

	client := &http.Client{}

	url := fmt.Sprintf("/token/%s", token)
	uri := fmt.Sprintf("https://%s%s", api_server, url)

	fmt.Println(uri)

	request, err := http.NewRequest("GET", uri, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	}

	signature := hmacSignature(url, "")

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token-wtf-key", api_key)
	request.Header.Add("token-wtf-signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	} else {
		responseBody, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(response.Status))
		fmt.Println(response.Header)

		return string(responseBody), nil
	}
}

func claimToken(token string) (string, error) {

	client := &http.Client{}

	url := fmt.Sprintf("/token/%s/claim", token)
	uri := fmt.Sprintf("https://%s%s", api_server, url)

	data := map[string]string{"api_key": api_key}
	json, _ := json2.Marshal(data)

	signature := hmacSignature(url, string(json))

	fmt.Println(uri)
	fmt.Println(string(json))

	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(json))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token-wtf-key", api_key)
	request.Header.Add("token-wtf-signature", signature)

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	} else {
		responseBody, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(response.Status))
		fmt.Println(response.Header)

		return string(responseBody), nil
	}
}

func getScore() (string, error) {

	uri := fmt.Sprintf("https://%s/score", api_server)

	fmt.Println(uri)

	response, err := http.Get(uri)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return "", err
	}

	fmt.Println(string(response.Status))

	return string(body), nil
}
