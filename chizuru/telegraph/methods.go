package telegraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Register(shortName string, authorName string, authorUrl string) (response *CreateAccount, err error) {
	/*
		Registers a user on telegra.ph
	*/
	reqURL := fmt.Sprintf("https://api.telegra.ph/createAccount?short_name=%s&author_name=%s&author_url=%s",
		url.QueryEscape(shortName), url.QueryEscape(authorName), url.QueryEscape(authorUrl))

	//fmt.Println(reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Printf("[telegraph][Register] Failed due to %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(CreateAccount)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
