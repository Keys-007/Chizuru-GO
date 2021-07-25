package nhentai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest(id int) (response *NHentaiResponse, err error) {
	reqURL := fmt.Sprintf("https://nhentai.net/api/gallery/%d", id)
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Printf("[nHentai][Request] Failed due to %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(NHentaiResponse)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
