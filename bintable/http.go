package bintable

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) []byte {
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	client := &http.Client{}

	body, err := client.Do(res)
	if err != nil {
		return nil
	}

	bd, err := ioutil.ReadAll(body.Body)
	if err != nil {
		return nil
	}

	return bd

}
