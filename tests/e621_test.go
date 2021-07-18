package tests

import (
	"Chizuru-GO/chizuru/e621API"
	"log"
	"testing"
)

func TestE621API(t *testing.T) {
	r, err := e621API.DoRequest("fox")
	if err != nil {
		log.Println(err.Error())
	}
	if r == nil {
		log.Println("Response is nil")
		return
	}
	log.Println(r.Posts[0])
}
