package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DoGetRequest(url, ref string) []byte {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}

	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Referer", ref)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return bodyBytes
}

func GetResponseHeaders(url, ref string) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}

	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Referer", ref)
	if err != nil {
		log.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	for i := range resp.Cookies() {
		log.Println(resp.Cookies()[i])
	}

	req, err = http.NewRequest("GET", TARGET_AUTH_CODE_URL, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Referer", ref)
	req.Header.Set("Cookie", fmt.Sprint(resp.Cookies()[0]))
	if err != nil {
		log.Println(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	bodybytes, _ := ioutil.ReadAll(resp.Body)

	log.Println("Status code: ", resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println(fmt.Sprint(bodybytes))
	} else {
		log.Println(string(bodybytes))
	}

}
