package e621API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest(tag string) (res *E621Response, err error) {
	reqUrl := fmt.Sprintf("https://e621.net/posts.json?tags=%s", tag)
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close() // ??? how am i supposed to handle the error here
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	data := new(E621Response)
	err = json.Unmarshal(d, &data)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return data, nil

}
